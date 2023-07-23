.PHONY: all
all:

.PHONY: test
test:
	go test ./...

.PHONY: deps
deps:
	go mod tidy
