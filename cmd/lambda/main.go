package main

import (
	"log"

	"github.com/apex/gateway"

	"github.com/cszczepaniak/sudoku-solver/pkg/rest"
)

func main() {
	if err := gateway.ListenAndServe(`:8080`, rest.NewServer()); err != nil {
		log.Fatal(err)
	}
}
