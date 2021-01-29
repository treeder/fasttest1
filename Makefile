.PHONY: dep build docker run release install test backup

build:
	go build -o app

run: build
	./app

test:
	go test ./...
