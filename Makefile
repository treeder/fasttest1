.PHONY: dep build docker run release install test backup

build:
	go build

run: build
	./fasttest1

test:
	go test ./...
