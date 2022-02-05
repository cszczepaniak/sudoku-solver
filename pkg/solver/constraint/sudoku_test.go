package constraint

import (
	"testing"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/model"
	"github.com/stretchr/testify/require"
)

func TestSudokuIsPointConstraint(t *testing.T) {
	s := NewSudoku()
	require.Error(t, s.Evaluate(123))
}

func TestSudokuEvaluate(t *testing.T) {
	s := NewSudoku()

	s.AddValue(1, model.NewPoint(0, 0))
	require.Error(t, s.EvaluateAt(1, model.NewPoint(0, 5)))
	require.Error(t, s.EvaluateAt(1, model.NewPoint(8, 0)))
	require.Error(t, s.EvaluateAt(1, model.NewPoint(2, 2)))

	s.RemoveValue(1, model.NewPoint(0, 0))
	require.NoError(t, s.EvaluateAt(1, model.NewPoint(0, 5)))
	require.NoError(t, s.EvaluateAt(1, model.NewPoint(8, 0)))
	require.NoError(t, s.EvaluateAt(1, model.NewPoint(2, 2)))
}
