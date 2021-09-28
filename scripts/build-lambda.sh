set -ex

go test ./...

GOOS=linux CGO_ENABLED=0 go build -o sudoku-lambda -tags=prod cmd/lambda/main.go
zip $GITHUB_SHA.zip sudoku-lambda
rm sudoku-lambda
