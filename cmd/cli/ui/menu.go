package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rivo/tview"
)

type menuItem struct {
	name string
	desc string
}

func initializeMenu(items []menuItem) *tview.TextView {
	var maxLen int
	for _, item := range items {
		if len(item.name) > maxLen {
			maxLen = len(item.name)
		}
	}
	sb := strings.Builder{}
	for i, item := range items {
		fmt.Fprintf(&sb, `%-`+strconv.Itoa(maxLen)+`s - %s`, item.name, item.desc)
		if i != len(items)-1 {
			fmt.Fprintln(&sb)
		}
	}
	txt := tview.NewTextView()
	txt.SetTitle(`Sudoku Solver`).SetBorder(true)
	txt.SetText(sb.String())
	return txt
}
