package constraint

import (
	"errors"
	"fmt"
)

type Constraint interface {
	// Evaluate takes a candidate value for this constraint and returns an error if the constraint would
	// be violated by the value, otherwise nil.
	Evaluate(n int) error

	AddValue(n int)
	RemoveValue(n int)
}

func violation(val int, err error) error {
	return fmt.Errorf(`value %d violates constraint: %w`, val, err)
}

type uniquenessConstraint struct {
	set map[int]struct{}
}

func NewUniqueness() *uniquenessConstraint {
	return &uniquenessConstraint{
		set: make(map[int]struct{}),
	}
}

var errDuplicateValue = errors.New(`duplicate value`)

func (c *uniquenessConstraint) Evaluate(n int) error {
	if _, ok := c.set[n]; ok {
		return violation(n, errDuplicateValue)
	}
	return nil
}

func (c *uniquenessConstraint) AddValue(n int) {
	c.set[n] = struct{}{}
}

func (c *uniquenessConstraint) RemoveValue(n int) {
	delete(c.set, n)
}
