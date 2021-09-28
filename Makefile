.PHONY: cli
cli:
	go run cmd/cli/main.go

.PHONY: rest
rest:
	go run cmd/rest/main.go