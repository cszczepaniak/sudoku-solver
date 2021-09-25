package solver

import (
	"errors"
)

var (
	ErrWrongNumberOfRows = errors.New(`expected 9 rows`)
	ErrWrongNumberOfCols = errors.New(`expected 9 cols`)
	ErrNumberOutOfRange  = errors.New(`expected numbers 0-9`)
	ErrInvalidBoard      = errors.New(`board was not a valid sudoku board`)
	ErrNoSolution        = errors.New(`no solution exists for the given board`)
)

const (
	Dimension    = 9
	TotalSquares = Dimension * Dimension
	Empty        = 0
	MinEntry     = 1
	MaxEntry     = 9
)

func NewEmptyBoard() [][]int {
	res := make([][]int, 9)
	for i := range res {
		res[i] = make([]int, 9)
	}
	return res
}

type Solver struct {
	nums  [TotalSquares]int
	rows  [Dimension][Dimension]int
	cols  [Dimension][Dimension]int
	boxes [Dimension][Dimension]int
}

func New(board [][]int) (*Solver, error) {
	if len(board) != Dimension {
		return nil, ErrWrongNumberOfRows
	}
	s := &Solver{}
	for i, r := range board {
		if len(r) != Dimension {
			return nil, ErrWrongNumberOfCols
		}
		for j, n := range r {
			if n < Empty || n > MaxEntry {
				return nil, ErrNumberOutOfRange
			}
			if n == 0 {
				continue
			}
			if !s.writeAtRowCol(n, i, j) {
				return nil, ErrInvalidBoard
			}
		}
	}
	return s, nil
}

func (s *Solver) ToBoard() [][]int {
	res := make([][]int, Dimension)
	for i := 0; i < Dimension; i++ {
		res[i] = s.nums[i*9 : (i*9)+9]
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
	if start >= len(s.nums) {
		return nil
	}
	for i, n := range s.nums[start:] {
		if n != 0 {
			// there's already a number here
			continue
		}
		for guess := MinEntry; guess <= MaxEntry; guess++ {
			if !s.isValidEntry(guess, start+i) {
				continue
			}
			s.writeAt(guess, i+start)
			if err := s.solveFrom(start + i + 1); err == ErrNoSolution {
				s.clearAt(guess, i+start)
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

func (s *Solver) isValidEntry(n, idx int) bool {
	return s.rows[rowFromIdx(idx)][n-1] == 0 &&
		s.cols[colFromIdx(idx)][n-1] == 0 &&
		s.boxes[boxFromIdx(idx)][n-1] == 0
}

func (s *Solver) writeAtRowCol(n, r, c int) bool {
	return s.writeAt(n, r*9+c)
}

// writeAt writes the given number at the given index and also updates internal puzzle state.
// It returns true if the given number was valid at the given index, and false otherwise
func (s *Solver) writeAt(n, idx int) bool {
	if !s.isValidEntry(n, idx) {
		return false
	}
	s.nums[idx] = n
	s.rows[rowFromIdx(idx)][n-1] = n
	s.cols[colFromIdx(idx)][n-1] = n
	s.boxes[boxFromIdx(idx)][n-1] = n
	return true
}

func (s *Solver) clearAt(n, idx int) {
	s.nums[idx] = 0
	s.rows[rowFromIdx(idx)][n-1] = 0
	s.cols[colFromIdx(idx)][n-1] = 0
	s.boxes[boxFromIdx(idx)][n-1] = 0
}

func rowFromIdx(idx int) int { return idx / 9 }
func colFromIdx(idx int) int { return idx % 9 }
func boxFromIdx(idx int) int { return 3*(idx/27) + (idx%9)/3 }
