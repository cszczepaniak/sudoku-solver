package ui

import (
	"fmt"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *Application) initializeTable() {
	tb := tview.NewTable().SetBorders(true)
	tb.Select(0, 0).SetSelectable(true, true).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			a.app.Stop()
		}
	}).SetSelectionChangedFunc(func(row, column int) {
		a.currRow, a.currCol = row, column
	}).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newVal := int(event.Rune() - '0')
			a.updateCell(a.currRow, a.currCol, newVal)
		case 'x':
			a.board = solver.NewEmptyBoard()
			a.redrawBoard()
		case 'e':
			a.board = [][]int{
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
			a.redrawBoard()
		case 'q':
			a.app.Stop()
		}

		if event.Key() == tcell.KeyEnter {
			s, err := solver.New(a.board, nil)
			if err == nil {
				solved, err := s.Solve()
				if err == nil {
					a.board = solved
					a.redrawBoard()
				}
			}
		}
		return event
	})
	a.table = tb
	a.redrawBoard()
}

func (a *Application) updateCell(r, c, n int) {
	a.board[r][c] = n
	a.redrawCell(r, c, n)
}

func (a *Application) redrawCell(r, c, n int) {
	str := `   `
	if n > 0 {
		str = fmt.Sprintf(` %d `, n)
	}
	a.table.SetCell(r, c, tview.NewTableCell(str).SetAlign(tview.AlignCenter))
}

func (a *Application) redrawBoard() {
	for i, r := range a.board {
		for j, n := range r {
			a.redrawCell(i, j, n)
		}
	}
}
