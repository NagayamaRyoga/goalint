.PHONY: all
all: build

.PHONY: build
build:
	go build -o goalint ./cmd/goalint

.PHONY: test
test:
	go test ./...

.PHONY: test-update-snaps
test-update-snaps:
	UPDATE_SNAPS=true go test ./...

.PHONY: lint
lint:
	golangci-lint run --allow-parallel-runners

.PHONY: deps
deps:
	go mod tidy

.PHONY: examples
examples:
	${MAKE} -C _examples
