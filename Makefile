dev:
	go run .

build:
	go build

test:
	go test -v ./...

deploy: test build
