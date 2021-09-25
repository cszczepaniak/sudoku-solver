package cli

import (
	"github.com/cszczepaniak/sudoku-solver/pkg/solver"
	"github.com/rivo/tview"
)

type Application struct {
	board   [][]int
	currRow int
	currCol int

	table *tview.Table
	app   *tview.Application
}

func New(board [][]int) *Application {
	if board == nil {
		board = solver.NewEmptyBoard()
	}
	a := &Application{
		app:     tview.NewApplication(),
		board:   board,
		currRow: 0,
		currCol: 0,
	}
	a.initializeTable()

	txt := initializeMenu([]menuItem{{
		name: `Arrow Keys`,
		desc: `Move Selection`,
	}, {
		name: `1-9`,
		desc: `Set Value`,
	}, {
		name: `X`,
		desc: `Clear Puzzle`,
	}, {
		name: `E`,
		desc: `Load Example Puzzle`,
	}, {
		name: `Enter`,
		desc: `Solve Puzzle`,
	}, {
		name: `Esc`,
		desc: `Quit`,
	}})

	f := tview.NewFlex().SetDirection(tview.FlexColumn)
	f.AddItem(txt, 0, 1, false)
	f.AddItem(a.table, 0, 2, true)
	a.app.SetRoot(f, true)

	return a
}

func (a *Application) Run() error {
	return a.app.Run()
}
