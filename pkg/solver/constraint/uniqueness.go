package constraint

import (
	"errors"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

type uniquenessConstraint struct {
	set map[int]map[model.Point]struct{}
}

func NewUniqueness() *uniquenessConstraint {
	return &uniquenessConstraint{
		set: make(map[int]map[model.Point]struct{}),
	}
}

var errDuplicateValue = errors.New(`duplicate value`)

func (c *uniquenessConstraint) Evaluate(n int) error {
	if pts, ok := c.set[n]; ok && len(pts) != 0 {
		return violation(n, errDuplicateValue)
	}
	return nil
}

func (c *uniquenessConstraint) Validate() error {
	var violations []model.Point
	for _, pts := range c.set {
		if len(pts) > 1 {
			for pt := range pts {
				violations = append(violations, pt)
			}
		}
	}
	if len(violations) > 0 {
		return &ValidationError{
			Points:  violations,
			Message: errDuplicateValue.Error(),
		}
	}
	return nil
}

func (c *uniquenessConstraint) AddValue(n int, p model.Point) {
	if _, ok := c.set[n]; !ok {
		c.set[n] = make(map[model.Point]struct{})
	}
	c.set[n][p] = struct{}{}
}

func (c *uniquenessConstraint) RemoveValue(n int, p model.Point) {
	delete(c.set[n], p)
}
