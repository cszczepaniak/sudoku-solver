package main

import (
	"log"

	"github.com/cszczepaniak/sudoku-solver/pkg/solver/cli"
)

func main() {
	if err := cli.New(nil).Run(); err != nil {
		log.Fatal(err)
	}
}
