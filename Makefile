.PHONY: build, release, lint

build:
	go build -o mu-worker cmd/main.go

release:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o mu-worker cmd/main.go

lint:
	golangci-lint run
