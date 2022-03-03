## help - Display help about make targets for this Makefile
help:
	@cat Makefile | grep '^## ' --color=never | cut -c4- | sed -e "`printf 's/ - /\t- /;'`" | column -s "`printf '\t'`" -t

## install - Install globally from source
install:
	go build -o $(go env GOPATH)/bin/tuneuptechnology-go

## clean - Clean the project
clean:
	rm dist
	rm $(go env GOPATH)/bin/tuneuptechnology-go

## build - Build the project
build:
	go build -o dist/tuneuptechnology-go

## test - Test the project
test:
	go test ./test

## coverage - Get test coverage
coverage: 
	go test ./test -coverprofile=covprofile -coverpkg=./... && go tool cover -html=covprofile

## lint - Lint the project
lint:
	golangci-lint run

.PHONY: help install clean build test lint
