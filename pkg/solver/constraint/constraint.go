package constraint

import (
	"errors"
	"fmt"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

type Constraint interface {
	// Evaluate takes a candidate value for this constraint and returns an error if the constraint would
	// be violated by the value, otherwise nil.
	Evaluate(val int) error

	// Validate returns an error indicating which points are violating the constraint. If the constraint
	// is satisfied, it returns nil.
	Validate() error

	// AddValue is used to register a value with the constraint at the given point.
	AddValue(val int, p model.Point)
	// RemoveValue is used to unregister a value with the constraint at the given point.
	RemoveValue(val int, p model.Point)
}

type PointConstraint interface {
	Constraint

	// EvaluateAt takes a candidate point for this constraint and returns an error if the constraint
	// would be violated by adding the value to the given point.
	EvaluateAt(val int, p model.Point) error
}

type pointConstraintBase struct{}

func (pointConstraintBase) Evaluate(val int) error {
	return errors.New(`point constraint must be evaluated at a point`)
}

func violation(val int, err error) error {
	return fmt.Errorf(`value %d violates constraint: %w`, val, err)
}
