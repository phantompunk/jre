VERSION := localdev

default: test

run:
	go run main.go

lint:
	golangci-lint run

test:
	go test ./...

test-race:
	go test -race ./...

build: 
	docker build -t jrerest:$(VERSION) .

up:
	docker-compose up

.PHONY: all build run test clean 

