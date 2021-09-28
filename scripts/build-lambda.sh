#!/bin/bash
set -ex

source ./scripts/variables.sh

go test ./...

GOOS=linux CGO_ENABLED=0 go build -o sudoku-lambda -tags=prod cmd/lambda/main.go
zip $BUILD_NAME.zip sudoku-lambda
rm sudoku-lambda
