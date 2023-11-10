SRC=src/main.go
CC = go

test:
	$(CC) test -coverpkg=./... ./...

build:
	$(CC) build
