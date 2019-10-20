all: build

build:
	protoc --micro_out=. --go_out=. ./*/*/*/*/*.proto
	go build ./...

install:
	go install ./...