# Makefile for lyanDB

build:
	go build -o lyanDB main.go resp.go

run: build
	./lyanDB

test:
	go test -v

clean:
	rm -f lyanDB