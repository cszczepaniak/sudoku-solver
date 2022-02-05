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
