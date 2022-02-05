package constraint

import (
	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

// Sudoku is a helper representing the constraints for classic sudoku.
type Sudoku struct {
	pointConstraintBase
	Rows  [model.Dimension]*uniquenessConstraint
	Cols  [model.Dimension]*uniquenessConstraint
	Boxes [model.Dimension]*uniquenessConstraint
}

var _ PointConstraint = (*Sudoku)(nil)

func NewSudoku() *Sudoku {
	s := Sudoku{}
	for i := 0; i < model.Dimension; i++ {
		s.Rows[i] = NewUniqueness()
		s.Cols[i] = NewUniqueness()
		s.Boxes[i] = NewUniqueness()
	}
	return &s
}

func (s *Sudoku) EvaluateAt(val int, p model.Point) error {
	if err := s.Rows[p.Row()].Evaluate(val); err != nil {
		return err
	}
	if err := s.Cols[p.Col()].Evaluate(val); err != nil {
		return err
	}
	if err := s.Boxes[p.Box()].Evaluate(val); err != nil {
		return err
	}
	return nil
}

func (s *Sudoku) Validate() error {
	violations := make(map[model.Point]struct{})
	add := func(pts []model.Point) {
		for _, p := range pts {
			violations[p] = struct{}{}
		}
	}
	for i := 0; i < model.Dimension; i++ {
		pts, err := extractViolations(s.Rows[i].Validate())
		if err != nil {
			return err
		}
		add(pts)

		pts, err = extractViolations(s.Cols[i].Validate())
		if err != nil {
			return err
		}
		add(pts)

		pts, err = extractViolations(s.Boxes[i].Validate())
		if err != nil {
			return err
		}
		add(pts)
	}

	if len(violations) == 0 {
		return nil
	}
	verr := &ValidationError{
		Message: `sudoku constraints violated`,
	}
	for pt := range violations {
		verr.Points = append(verr.Points, pt)
	}
	return verr
}

func extractViolations(err error) ([]model.Point, error) {
	if err == nil {
		return nil, nil
	}

	verr, ok := err.(*ValidationError)
	if !ok {
		return nil, err
	}

	return verr.Points, nil
}

func (s *Sudoku) AddValue(val int, p model.Point) {
	s.Rows[p.Row()].AddValue(val, p)
	s.Cols[p.Col()].AddValue(val, p)
	s.Boxes[p.Box()].AddValue(val, p)
}

func (s *Sudoku) RemoveValue(val int, p model.Point) {
	s.Rows[p.Row()].RemoveValue(val, p)
	s.Cols[p.Col()].RemoveValue(val, p)
	s.Boxes[p.Box()].RemoveValue(val, p)
}
