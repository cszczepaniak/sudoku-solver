package solver

import (
	"errors"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/constraint"
	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

const (
	Dimension    = 9
	TotalSquares = Dimension * Dimension
	Empty        = 0
	MinEntry     = 1
	MaxEntry     = 9
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
	cells [Dimension][Dimension]*Cell
}

func New(board [][]int) (*Solver, error) {
	if len(board) != Dimension {
		return nil, ErrWrongNumberOfRows
	}

	s := &Solver{}
	sudokuConstraints := constraint.NewSudoku()
	for i, r := range board {
		if len(r) != Dimension {
			return nil, ErrWrongNumberOfCols
		}
		for j, n := range r {
			p := model.NewPoint(i, j)
			s.cells[i][j] = NewCell(
				p,
				sudokuConstraints,
			)
			if n == Empty {
				continue
			}
			// write without regard to duplicates or out of bounds; we'll validate those in a sec
			s.cells[i][j].Write(n)
		}
	}
	aggErr := constraint.NewAggregateValidationError()
	for _, r := range s.cells {
		for _, c := range r {
			if err := aggErr.Add(c.Validate()); err != nil {
				return nil, err
			}
		}
	}
	if err := aggErr.ToValidationError(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Solver) ToBoard() [][]int {
	res := make([][]int, Dimension)
	for i := 0; i < Dimension; i++ {
		res[i] = make([]int, Dimension)
		for j := 0; j < Dimension; j++ {
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
		for guess := MinEntry; guess <= MaxEntry; guess++ {
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
