package main

import (
	"log"

	"github.com/cszczepaniak/sudoku-solver/cmd/cli/ui"
)

func main() {
	if err := ui.NewApp(nil).Run(); err != nil {
		log.Fatal(err)
	}
}
