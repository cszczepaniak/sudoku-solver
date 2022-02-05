package solver

import (
	"github.com/cszczepaniak/sudoku-solver/pkg/solver/constraint"
	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
)

type Cell struct {
	Location    model.Point
	Value       int
	Constraints []constraint.Constraint
}

func NewCell(loc model.Point, constraints ...constraint.Constraint) *Cell {
	return &Cell{
		Location:    loc,
		Constraints: constraints,
	}
}

func (c *Cell) Write(val int) {
	for _, constraint := range c.Constraints {
		constraint.AddValue(val, c.Location)
	}
	c.Value = val
}

func (c *Cell) SatisfiesConstraints(val int) bool {
	for _, con := range c.Constraints {
		var err error
		switch t := con.(type) {
		case constraint.PointConstraint:
			err = t.EvaluateAt(val, c.Location)
		case constraint.Constraint:
			err = t.Evaluate(val)
		}
		if err != nil {
			return false
		}
	}
	return true
}

func (c *Cell) Clear() {
	for _, constraint := range c.Constraints {
		constraint.RemoveValue(c.Value, c.Location)
	}
	c.Value = model.Empty
}
