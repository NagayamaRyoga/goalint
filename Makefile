.PHONY: all
all:

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: deps
deps:
	go mod tidy
