.PHONY: build test generate run

generate:
	go generate ./...

setup:
	go mod tidy

build:
	go build ./...

test:
	go test ./...

run:
	go run ./cmd/server
