package solver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPoint(t *testing.T) {
	tests := []struct {
		row      int
		col      int
		expBox   int
		expIndex int
	}{{
		row:      0,
		col:      0,
		expBox:   0,
		expIndex: 0,
	}, {
		row:      1,
		col:      5,
		expBox:   1,
		expIndex: 14,
	}, {
		row:      2,
		col:      6,
		expBox:   2,
		expIndex: 24,
	}, {
		row:      5,
		col:      1,
		expBox:   3,
		expIndex: 46,
	}, {
		row:      4,
		col:      4,
		expBox:   4,
		expIndex: 40,
	}, {
		row:      3,
		col:      8,
		expBox:   5,
		expIndex: 35,
	}, {
		row:      8,
		col:      0,
		expBox:   6,
		expIndex: 72,
	}, {
		row:      7,
		col:      3,
		expBox:   7,
		expIndex: 66,
	}, {
		row:      8,
		col:      6,
		expBox:   8,
		expIndex: 78,
	}}
	for _, tc := range tests {
		pt := newPoint(tc.row, tc.col)
		require.Equal(t, tc.row, pt.row)
		require.Equal(t, tc.col, pt.col)
		require.Equal(t, tc.expBox, pt.box)
	}
}
