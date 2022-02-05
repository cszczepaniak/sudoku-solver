package constraint

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniquenessConstraint(t *testing.T) {
	c := NewUniqueness()
	for i := 1; i < 10; i++ {
		require.NoError(t, c.Evaluate(i))
	}

	c.AddValue(1)
	c.AddValue(2)
	c.AddValue(3)
	for i := 1; i < 4; i++ {
		err := c.Evaluate(i)
		require.Error(t, err)
		require.True(t, errors.Is(err, errDuplicateValue))
	}
	for i := 4; i < 10; i++ {
		require.NoError(t, c.Evaluate(i))
	}

	c.RemoveValue(2)
	require.NoError(t, c.Evaluate(2))
}
