.PHONY: cli
cli:
	go run cmd/cli/main.go

.PHONY: rest
rest:
	go run cmd/rest/main.go

.PHONY: lambda
lambda:
	GOOS=linux CGO_ENABLED=0 go build -o sudoku-lambda -tags=prod cmd/lambda/main.go
	zip sudoku-lambda.zip sudoku-lambda
	rm sudoku-lambda