package constraint

import (
	"fmt"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

// killerConstraint is an implementation of a constraint which requires a certain cage of digits to
// sum to a given number.
type killerConstraint struct {
	pointConstraintBase
	target    int
	cageCells map[model.Point]int
}

var _ PointConstraint = (*killerConstraint)(nil)

func NewKiller(target int, points ...model.Point) *killerConstraint {
	res := &killerConstraint{
		target:    target,
		cageCells: make(map[model.Point]int, len(points)),
	}
	for _, pt := range points {
		res.cageCells[pt] = 0
	}
	return res
}

func (c *killerConstraint) current() (int, int) {
	var res int
	var nEmpty int
	for _, val := range c.cageCells {
		res += val
		if val == 0 {
			nEmpty++
		}
	}
	return res, nEmpty
}

func (c *killerConstraint) EvaluateAt(val int, p model.Point) error {
	if _, ok := c.cageCells[p]; !ok {
		// this point has no bearing on the constraint
		return nil
	}
	curr, nEmpty := c.current()

	currCellVal := c.cageCells[p]
	newSum := curr - currCellVal + val

	if nEmpty <= 1 && newSum != c.target {
		// this would fill the cage with an incorrect sum
		return errWrongCageSum(c.target, newSum)
	} else if nEmpty > 1 && newSum >= c.target {
		// this wouldn't fill the box, but we're already over
		return errWrongCageSum(c.target, newSum)
	}

	return nil
}

func (c *killerConstraint) Validate() error {
	curr, nEmpty := c.current()
	if nEmpty >= 1 && curr < c.target {
		// our cage isn't full and we're below our target; that's fine
		return nil
	}
	if nEmpty == 0 && curr == c.target {
		// we're full and at the correct sum
		return nil
	}
	verr := &ValidationError{
		Message: errWrongCageSum(c.target, curr).Error(),
	}
	for pt := range c.cageCells {
		verr.Points = append(verr.Points, pt)
	}
	return verr
}

func errWrongCageSum(want, got int) error {
	return fmt.Errorf(`cage sum was %d; wanted %d`, got, want)
}

func (c *killerConstraint) AddValue(val int, p model.Point) {
	if _, ok := c.cageCells[p]; !ok {
		// this point isn't in the cage
		return
	}
	c.cageCells[p] = val
}

func (c *killerConstraint) RemoveValue(val int, p model.Point) {
	if _, ok := c.cageCells[p]; !ok {
		// this point isn't in the cage
		return
	}
	c.cageCells[p] = 0
}
