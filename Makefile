.PHONY: all
all: build

.PHONY: build
build:
	go build -o goalint ./cmd/goalint

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --allow-parallel-runners

.PHONY: deps
deps:
	go mod tidy
