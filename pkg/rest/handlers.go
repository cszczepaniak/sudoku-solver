package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver"
)

func (s *Server) solve(c *gin.Context) {
	input, ok := bindSudokuBoard(c)
	if !ok {
		return
	}
	solver, err := solver.New(input)
	if err != nil {
		writeErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	solution, err := solver.Solve()
	if err != nil {
		writeErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, solution)
}

func bindSudokuBoard(c *gin.Context) ([][]int, bool) {
	var input [][]int
	if err := c.BindJSON(&input); err != nil {
		writeErrorResponse(c, http.StatusBadRequest, err)
		return nil, false
	}
	return input, true
}

func writeErrorResponse(c *gin.Context, code int, err error) {
	resp := gin.H{
		`error`: err.Error(),
	}
	switch terr := err.(type) {
	case *solver.InvalidBoardError:
		resp[`invalidSquares`] = terr.InvalidSquares
	default:
	}
	c.JSON(code, resp)
}
