package main

import (
	"log"
	"net/http"

	"github.com/cszczepaniak/sudoku-solver/pkg/rest"
)

func main() {
	if err := http.ListenAndServe(`:8080`, rest.NewServer()); err != nil {
		log.Fatal(err)
	}
}
