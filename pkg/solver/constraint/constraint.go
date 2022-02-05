package constraint

import (
	"errors"
	"fmt"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

type ValidationError struct {
	Points  []model.Point
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type Constraint interface {
	// Evaluate takes a candidate value for this constraint and returns an error if the constraint would
	// be violated by the value, otherwise nil.
	Evaluate(n int) error

	// Validate returns an error indicating which points are violating the constraint. If the constraint
	// is satisfied, it returns nil.
	Validate() error

	// AddValue is used to register a value with the constraint at the given point.
	AddValue(n int, p model.Point)
	// RemoveValue is used to unregister a value with the constraint at the given point.
	RemoveValue(n int, p model.Point)
}

func violation(val int, err error) error {
	return fmt.Errorf(`value %d violates constraint: %w`, val, err)
}

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
	for _, pts := range c.set {
		if len(pts) > 1 {
			ptSlice := make([]model.Point, 0, len(pts))
			for pt := range pts {
				ptSlice = append(ptSlice, pt)
			}
			return &ValidationError{
				Points:  ptSlice,
				Message: errDuplicateValue.Error(),
			}
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
