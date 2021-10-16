package solver_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver"
)

func TestSolve(t *testing.T) {
	input := [][]int{
		{0, 0, 9, 0, 1, 6, 0, 4, 2},
		{1, 0, 4, 2, 0, 9, 0, 6, 0},
		{0, 2, 0, 0, 0, 8, 7, 0, 0},
		{3, 5, 0, 0, 9, 0, 1, 0, 0},
		{0, 6, 7, 4, 0, 1, 9, 0, 5},
		{0, 0, 0, 7, 5, 0, 0, 8, 6},
		{0, 9, 0, 0, 0, 4, 8, 5, 7},
		{8, 0, 0, 9, 6, 0, 0, 2, 0},
		{4, 7, 0, 8, 0, 5, 0, 0, 0},
	}
	solved := [][]int{
		{7, 8, 9, 5, 1, 6, 3, 4, 2},
		{1, 3, 4, 2, 7, 9, 5, 6, 8},
		{5, 2, 6, 3, 4, 8, 7, 1, 9},
		{3, 5, 8, 6, 9, 2, 1, 7, 4},
		{2, 6, 7, 4, 8, 1, 9, 3, 5},
		{9, 4, 1, 7, 5, 3, 2, 8, 6},
		{6, 9, 2, 1, 3, 4, 8, 5, 7},
		{8, 1, 5, 9, 6, 7, 4, 2, 3},
		{4, 7, 3, 8, 2, 5, 6, 9, 1},
	}
	s, err := solver.New(input)
	require.NoError(t, err)

	actual, err := s.Solve()
	require.NoError(t, err)
	require.Equal(t, solved, actual)
}

func TestNoSolution(t *testing.T) {
	input := [][]int{
		{5, 1, 6, 8, 4, 9, 7, 3, 2},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 5},
		{1, 3, 5, 0, 6, 0, 9, 0, 7},
		{4, 7, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 7, 4},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 8},
	}
	s, err := solver.New(input)
	require.NoError(t, err)

	actual, err := s.Solve()
	require.Equal(t, solver.ErrNoSolution, err)
	require.Nil(t, actual)
}

func TestToBoard(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}
	s, err := solver.New(input)
	require.NoError(t, err)
	require.Equal(t, input, s.ToBoard())
}

func TestNew(t *testing.T) {
	tests := []struct {
		desc   string
		input  [][]int
		expErr error
	}{{
		desc:   `not enough rows`,
		input:  [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		expErr: solver.ErrWrongNumberOfRows,
	}, {
		desc: `too many rows`,
		input: [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		expErr: solver.ErrWrongNumberOfRows,
	}, {
		desc:   `not enough cols`,
		input:  [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
		expErr: solver.ErrWrongNumberOfCols,
	}, {
		desc: `too many cols`,
		input: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: solver.ErrWrongNumberOfCols,
	}, {
		desc: `number too big`,
		input: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 10, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 7,
				Col: 6,
				Msg: `number out of range`,
			}},
		},
	}, {
		desc: `number too small`,
		input: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, -2, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 2,
				Col: 2,
				Msg: `number out of range`,
			}},
		},
	}, {
		desc: `duplicate in row`,
		input: [][]int{
			{1, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 0,
				Col: 0,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 0,
				Col: 3,
				Msg: `duplicate number in row, column, or box`,
			}},
		},
	}, {
		desc: `duplicate in row and box`,
		input: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 1,
				Col: 0,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 1,
				Col: 2,
				Msg: `duplicate number in row, column, or box`,
			}},
		},
	}, {
		desc: `duplicate in col`,
		input: [][]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 0,
				Col: 0,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 4,
				Col: 0,
				Msg: `duplicate number in row, column, or box`,
			}},
		},
	}, {
		desc: `duplicate in col and box`,
		input: [][]int{
			{0, 0, 0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 0,
				Col: 5,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 1,
				Col: 5,
				Msg: `duplicate number in row, column, or box`,
			}},
		},
	}, {
		desc: `duplicate in box`,
		input: [][]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 0,
				Col: 0,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 2,
				Col: 2,
				Msg: `duplicate number in row, column, or box`,
			}},
		},
	}, {
		desc: `a lot of duplicates`,
		input: [][]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 3, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1},
			{0, 0, 1, 0, 0, 2, 0, 0, 0},
		},
		expErr: &solver.InvalidBoardError{
			Errors: []*solver.InvalidSquareError{{
				Row: 0,
				Col: 0,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 0,
				Col: 8,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 2,
				Col: 2,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 7,
				Col: 8,
				Msg: `duplicate number in row, column, or box`,
			}, {
				Row: 8,
				Col: 2,
				Msg: `duplicate number in row, column, or box`,
			}},
		},
	}, {
		desc: `no error`,
		input: [][]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 2, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 3, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 4, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 5, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 6, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 7, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 8, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 9},
		},
		expErr: nil,
	}}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			_, err := solver.New(tc.input)
			switch exp := tc.expErr.(type) {
			case *solver.InvalidBoardError:
				require.IsType(t, &solver.InvalidBoardError{}, err)
				require.ElementsMatch(t, exp.Errors, err.(*solver.InvalidBoardError).Errors)
			default:
				require.Equal(t, tc.expErr, err)
			}
		})
	}
}
