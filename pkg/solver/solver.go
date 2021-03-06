package solver

import "errors"

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
	nums  [TotalSquares]int
	cache *puzzleCache
}

func New(board [][]int) (*Solver, error) {
	if len(board) != Dimension {
		return nil, ErrWrongNumberOfRows
	}
	s := &Solver{
		cache: newPuzzleCache(),
	}
	var errs []*InvalidSquareError
	for i, r := range board {
		if len(r) != Dimension {
			return nil, ErrWrongNumberOfCols
		}
		for j, n := range r {
			if n < Empty || n > MaxEntry {
				errs = append(errs, newInvalidSquareError(i, j, outOfRange))
				continue
			}
			if n == 0 {
				continue
			}
			// write without regard to duplicates; we'll validate those later
			s.writeAt(i, j, n)
		}
	}

	errs = append(errs, s.cache.validateDuplicates()...)
	if len(errs) != 0 {
		return nil, &InvalidBoardError{
			InvalidSquares: errs,
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
			idx := start + i
			r, c := idx/9, idx%9
			if !s.cache.isValidEntry(r, c, guess) {
				continue
			}
			s.writeAt(r, c, guess)
			if err := s.solveFrom(start + i + 1); err == ErrNoSolution {
				s.clearAt(r, c, guess)
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

func (s *Solver) writeAt(r, c, n int) {
	s.nums[r*9+c] = n
	s.cache.add(r, c, n)
}

func (s *Solver) clearAt(r, c, n int) {
	s.nums[r*9+c] = 0
	s.cache.remove(r, c, n)
}
