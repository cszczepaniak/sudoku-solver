package constraint

import (
	"errors"
	"testing"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
	"github.com/stretchr/testify/require"
)

func TestUniquenessEvaluate(t *testing.T) {
	c := NewUniqueness()
	for i := 1; i < 10; i++ {
		require.NoError(t, c.Evaluate(i))
	}

	c.AddValue(1, model.Point{})
	c.AddValue(2, model.Point{})
	c.AddValue(3, model.Point{})
	for i := 1; i < 4; i++ {
		err := c.Evaluate(i)
		require.Error(t, err)
		require.True(t, errors.Is(err, errDuplicateValue))
	}
	for i := 4; i < 10; i++ {
		require.NoError(t, c.Evaluate(i))
	}

	c.RemoveValue(2, model.Point{})
	require.NoError(t, c.Evaluate(2))
}

func TestUniquenessValidate(t *testing.T) {
	c := NewUniqueness()
	c.AddValue(1, model.NewPoint(1, 2))
	c.AddValue(2, model.NewPoint(2, 2))

	require.NoError(t, c.Validate())

	c.AddValue(1, model.NewPoint(2, 3))
	err := c.Validate()
	require.Error(t, err)
	require.IsType(t, &ValidationError{}, err)

	verr := err.(*ValidationError)
	require.ElementsMatch(t, []model.Point{
		model.NewPoint(1, 2),
		model.NewPoint(2, 3),
	}, verr.Points)

	c.AddValue(2, model.NewPoint(5, 5))
	c.AddValue(2, model.NewPoint(5, 6))
	c.AddValue(2, model.NewPoint(5, 7))
	err = c.Validate()
	require.Error(t, err)
	require.IsType(t, &ValidationError{}, err)

	verr = err.(*ValidationError)
	require.ElementsMatch(t, []model.Point{
		model.NewPoint(1, 2),
		model.NewPoint(2, 3),
		model.NewPoint(2, 2),
		model.NewPoint(5, 5),
		model.NewPoint(5, 6),
		model.NewPoint(5, 7),
	}, verr.Points)
}
