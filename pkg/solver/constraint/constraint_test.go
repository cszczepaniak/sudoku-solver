package constraint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSudokuConstraint(t *testing.T) {
	row := map[int]struct{}{}
	col := map[int]struct{}{}
	box := map[int]struct{}{}

	c := &sudokuConstraint{
		row: row,
		col: col,
		box: box,
	}

	for i := 1; i < 10; i++ {
		require.True(t, c.Evaluate(i))
	}

	row[1] = struct{}{}
	col[2] = struct{}{}
	box[3] = struct{}{}
	for i := 1; i < 4; i++ {
		require.False(t, c.Evaluate(i))
	}
	for i := 4; i < 10; i++ {
		require.True(t, c.Evaluate(i))
	}
}
