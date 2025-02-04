# Makefile for lyanDB

build:
	go build -o lyanDB main.go resp.go commands.go server.go storage.go

run: build
	./lyanDB

fmt:
	go fmt ./...

lint:
	go vet ./...

clean:
	rm -f lyanDB
	rm -rf *.test
