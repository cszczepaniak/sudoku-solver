package constraint

import "github.com/cszczepaniak/sudoku-solver/pkg/solver/model"

// Sudoku is a helper representing the constraints for classic sudoku.
type Sudoku struct {
	Rows  [model.Dimension]*uniquenessConstraint
	Cols  [model.Dimension]*uniquenessConstraint
	Boxes [model.Dimension]*uniquenessConstraint
}

func NewSudoku() Sudoku {
	s := Sudoku{}
	for i := 0; i < model.Dimension; i++ {
		s.Rows[i] = NewUniqueness()
		s.Cols[i] = NewUniqueness()
		s.Boxes[i] = NewUniqueness()
	}
	return s
}

func (s Sudoku) ConstraintsForPoint(p model.Point) []Constraint {
	return []Constraint{
		s.Rows[p.Row()],
		s.Cols[p.Col()],
		s.Boxes[p.Box()],
	}
}
