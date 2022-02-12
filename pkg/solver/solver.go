package solver

import (
	"errors"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/constraint"
	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

var (
	ErrWrongNumberOfRows = errors.New(`expected 9 rows`)
	ErrWrongNumberOfCols = errors.New(`expected 9 cols`)
	ErrNoSolution        = errors.New(`no solution exists for the given board`)
)

func NewEmptyBoard() [][]int {
	res := make([][]int, 9)
	for i := range res {
		res[i] = make([]int, 9)
	}
	return res
}

type Solver struct {
	cells [model.Dimension][model.Dimension]*Cell
}

func New(board [][]int, cellConstraints map[model.Point][]constraint.Constraint) (*Solver, error) {
	if len(board) != model.Dimension {
		return nil, ErrWrongNumberOfRows
	}

	s := &Solver{}
	sudokuConstraints := constraint.NewSudoku()
	for i, r := range board {
		if len(r) != model.Dimension {
			return nil, ErrWrongNumberOfCols
		}
		for j, n := range r {
			p := model.NewPoint(i, j)

			var cons []constraint.Constraint
			if cellConstraints == nil {
				cons = []constraint.Constraint{sudokuConstraints}
			} else {
				cons = make([]constraint.Constraint, 0, len(cellConstraints[p])+1)
				cons = append(cons, cellConstraints[p]...)
				cons = append(cons, sudokuConstraints)
			}

			s.cells[i][j] = NewCell(
				p,
				cons...,
			)
			if n == model.Empty {
				continue
			}
			// write without regard to duplicates or out of bounds; we'll validate those in a sec
			s.cells[i][j].Write(n)
		}
	}

	err := s.validateCells()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Solver) validateCells() error {
	aggErr := constraint.NewAggregateValidationError()
	for _, row := range s.cells {
		for _, cell := range row {
			for _, con := range cell.Constraints {
				if err := aggErr.Add(con.Validate()); err != nil {
					// This wasn't a validation error
					return err
				}
			}
		}
	}
	return aggErr.ToValidationError()
}

func (s *Solver) ToBoard() [][]int {
	res := make([][]int, model.Dimension)
	for i := 0; i < model.Dimension; i++ {
		res[i] = make([]int, model.Dimension)
		for j := 0; j < model.Dimension; j++ {
			res[i][j] = s.cells[i][j].Value
		}
	}
	return res
}

func (s *Solver) Solve() ([][]int, error) {
	if err := s.solveFrom(0); err != nil {
		return nil, err
	}
	return s.ToBoard(), nil
}

func (s *Solver) solveFrom(start int) error {
	for i := start; i < model.TotalSquares; i++ {
		r, c := i/9, i%9
		if s.cells[r][c].Value != 0 {
			// there's already a number here
			continue
		}
		for guess := model.MinEntry; guess <= model.MaxEntry; guess++ {
			if !s.cells[r][c].SatisfiesConstraints(guess) {
				continue
			}
			s.cells[r][c].Write(guess)
			if err := s.solveFrom(i + 1); err == ErrNoSolution {
				s.cells[r][c].Clear()
				continue
			} else if err != nil {
				return err
			}
			return nil
		}
		return ErrNoSolution
	}
	return nil
}
