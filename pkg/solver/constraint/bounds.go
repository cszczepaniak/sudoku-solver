package constraint

import (
	"errors"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

var errOutOfBounds = errors.New(`number out of range`)

const (
	min = 1
	max = 9
)

type boundsConstraint struct {
	vals map[model.Point]int
}

var _ Constraint = (*boundsConstraint)(nil)

func NewBounds() *boundsConstraint {
	return &boundsConstraint{
		make(map[model.Point]int),
	}
}

func (c *boundsConstraint) Evaluate(val int) error {
	if val == model.Empty {
		return nil
	}
	if val < min || val > max {
		return errOutOfBounds
	}
	return nil
}

func (c *boundsConstraint) Validate() error {
	var pts []model.Point
	for pt, val := range c.vals {
		if val == model.Empty {
			continue
		}
		if val < min || val > max {
			pts = append(pts, pt)
		}
	}

	if len(pts) == 0 {
		return nil
	}

	return &ValidationError{
		Message: errOutOfBounds.Error(),
		Points:  pts,
	}
}

func (c *boundsConstraint) AddValue(val int, pt model.Point) {
	c.vals[pt] = val
}

func (c *boundsConstraint) RemoveValue(_ int, pt model.Point) {
	delete(c.vals, pt)
}
