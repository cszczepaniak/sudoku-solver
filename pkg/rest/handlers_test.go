package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHealthHandler(t *testing.T) {
	ts := httptest.NewServer(NewServer())
	defer ts.Close()

	res, err := http.Get(ts.URL + `/api/health`)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Equal(t, `Healthy!`, string(readBytes(t, res.Body)))
}

func TestSolveSimpleErrors(t *testing.T) {
	ts := httptest.NewServer(NewServer())
	defer ts.Close()
	url := ts.URL + `/api/solve`

	res, err := http.Post(url, `application/json`, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	res.Body.Close()

	tests := []struct {
		desc       string
		board      [][]int
		getExpData func() gin.H
	}{{
		desc:       `empty board`,
		board:      [][]int{},
		getExpData: func() gin.H { return gin.H{`error`: solver.ErrWrongNumberOfRows.Error()} },
	}, {
		desc:       `wrong number of cols`,
		board:      [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
		getExpData: func() gin.H { return gin.H{`error`: solver.ErrWrongNumberOfCols.Error()} },
	}, {
		desc: `number out of range`,
		board: [][]int{
			{10, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		getExpData: func() gin.H {
			expErr := &solver.InvalidBoardError{
				Errors: []*solver.InvalidSquareError{{
					Row: 0,
					Col: 0,
					Msg: `number out of range`,
				}},
			}
			return gin.H{`error`: expErr.Error(), `detail`: expErr.Errors}
		},
	}, {
		desc: `no solution`,
		board: [][]int{
			{5, 1, 6, 8, 4, 9, 7, 3, 2},
			{3, 0, 7, 6, 0, 5, 0, 0, 0},
			{8, 0, 9, 7, 0, 0, 0, 6, 5},
			{1, 3, 5, 0, 6, 0, 9, 0, 7},
			{4, 7, 2, 5, 9, 1, 0, 0, 6},
			{9, 6, 8, 3, 7, 0, 0, 5, 0},
			{2, 5, 3, 1, 8, 6, 0, 7, 4},
			{6, 8, 4, 2, 0, 7, 5, 0, 0},
			{7, 9, 1, 0, 5, 0, 6, 0, 8},
		},
		getExpData: func() gin.H { return gin.H{`error`: solver.ErrNoSolution.Error()} },
	}}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			res, err := http.Post(url, `application/json`, boardToReader(t, tc.board))
			require.NoError(t, err)
			require.Equal(t, http.StatusBadRequest, res.StatusCode)
			compareResponse(t, tc.getExpData(), res.Body)
		})
	}
}

func TestSolveDuplicateSquareErrors(t *testing.T) {
	ts := httptest.NewServer(NewServer())
	defer ts.Close()
	url := ts.URL + `/api/solve`

	board := [][]int{
		{1, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	res, err := http.Post(url, `application/json`, boardToReader(t, board))
	require.NoError(t, err)

	expErr := &solver.InvalidBoardError{
		Errors: []*solver.InvalidSquareError{{
			Row: 0,
			Col: 0,
			Msg: `duplicate number in row, column, or box`,
		}, {
			Row: 0,
			Col: 3,
			Msg: `duplicate number in row, column, or box`,
		}, {
			Row: 2,
			Col: 0,
			Msg: `duplicate number in row, column, or box`,
		}},
	}

	fs := parseResponse(t, res.Body)

	errField, ok := fs[`error`]
	require.True(t, ok)
	require.Equal(t, expErr.Error(), errField)

	detailField, ok := fs[`detail`]
	require.True(t, ok)
	require.IsType(t, []interface{}{}, detailField)

	// we have to remarshal and unmarshal to get the strong typing we want
	bs, err := json.Marshal(detailField)
	require.NoError(t, err)
	var errs []*solver.InvalidSquareError
	require.NoError(t, json.Unmarshal(bs, &errs))

	require.Len(t, errs, len(expErr.Errors))
	for _, err := range errs {
		require.Contains(t, expErr.Errors, err)
	}
}

func TestSolve(t *testing.T) {
	ts := httptest.NewServer(NewServer())
	defer ts.Close()
	url := ts.URL + `/api/solve`

	board := [][]int{
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
	res, err := http.Post(url, `application/json`, boardToReader(t, board))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	compareResponse(t, solved, res.Body)
}

func boardToReader(t *testing.T, b [][]int) io.Reader {
	bs, err := json.Marshal(b)
	require.NoError(t, err)
	return bytes.NewReader(bs)
}

func readBytes(t *testing.T, r io.ReadCloser) []byte {
	defer r.Close()
	bs, err := ioutil.ReadAll(r)
	require.NoError(t, err)
	return bs
}

func compareResponse(t *testing.T, exp interface{}, body io.ReadCloser) {
	actual := readBytes(t, body)
	expBytes, err := json.Marshal(exp)
	require.NoError(t, err)
	require.Equal(t, expBytes, actual)
}

func parseResponse(t *testing.T, body io.ReadCloser) map[string]interface{} {
	actual := readBytes(t, body)
	var parsed map[string]interface{}
	require.NoError(t, json.Unmarshal(actual, &parsed))
	return parsed
}
