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

func (c *Cell) Write(val int) error {
	for _, constraint := range c.Constraints {
		err := constraint.Evaluate(val)
		if err != nil {
			return err
		}
	}
	for _, constraint := range c.Constraints {
		constraint.AddValue(val, c.Location)
	}
	c.Value = val
	return nil
}

func (c *Cell) Clear() {
	for _, constraint := range c.Constraints {
		constraint.RemoveValue(c.Value, c.Location)
	}
	c.Value = model.Empty
}

func (c *Cell) LinearIndex() int {
	return c.Location.LinearIndex()
}
