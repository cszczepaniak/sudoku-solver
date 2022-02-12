package constraint

import (
	"testing"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
	"github.com/stretchr/testify/require"
)

func TestKillerEvaluate(t *testing.T) {
	p1, p2, p3 := model.NewPoint(0, 0), model.NewPoint(1, 0), model.NewPoint(2, 0)
	c := NewKiller(15, map[model.Point]struct{}{p1: {}, p2: {}, p3: {}})

	// This is a PointConstraint
	require.Error(t, c.Evaluate(123))

	require.NoError(t, c.EvaluateAt(1, p1))

	c.AddValue(9, p1)
	require.Error(t, c.EvaluateAt(8, p2))
	require.NoError(t, c.EvaluateAt(8, model.NewPoint(8, 8)))

	c.AddValue(1, p2)
	for _, n := range []int{2, 3, 4, 6, 7, 8, 9} {
		require.Error(t, c.EvaluateAt(n, p3))
	}
	require.NoError(t, c.EvaluateAt(5, p3))

	c.AddValue(2, p3)
	require.Error(t, c.EvaluateAt(3, p3))
	require.NoError(t, c.EvaluateAt(5, p3))

	c.RemoveValue(1, p2)
	c.RemoveValue(2, p3)
	c.AddValue(5, p2)

	require.NoError(t, c.EvaluateAt(1, p3))
	for i := 2; i < 10; i++ {
		require.Error(t, c.EvaluateAt(i, p3))
	}
}

func TestKillerValidate(t *testing.T) {
	p1 := model.NewPoint(0, 0)
	p2 := model.NewPoint(1, 0)
	p3 := model.NewPoint(2, 0)
	p4 := model.NewPoint(3, 0)
	pts := map[model.Point]struct{}{
		p1: {},
		p2: {},
		p3: {},
		p4: {},
	}
	c := NewKiller(10, pts)

	require.NoError(t, c.Validate())

	c.AddValue(4, p1)
	c.AddValue(6, p2)
	err := c.Validate()
	require.Error(t, err)
	require.IsType(t, &ValidationError{}, err)
	require.ElementsMatch(t, pts, err.(*ValidationError).Points)

	c.RemoveValue(6, p2)
	c.AddValue(9, p2)
	err = c.Validate()
	require.Error(t, err)
	require.IsType(t, &ValidationError{}, err)
	require.ElementsMatch(t, pts, err.(*ValidationError).Points)

	c.RemoveValue(9, p2)
	c.AddValue(1, p2)
	c.AddValue(2, p3)
	err = c.Validate()
	require.NoError(t, err)

	for _, n := range []int{5, 6, 7, 8, 9} {
		c.AddValue(n, p4)

		err := c.Validate()
		require.Error(t, err)
		require.IsType(t, &ValidationError{}, err)
		require.ElementsMatch(t, pts, err.(*ValidationError).Points)

		c.RemoveValue(n, p4)
	}

	c.AddValue(3, p4)
	require.NoError(t, c.Validate())
}
