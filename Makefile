usage:				## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

all: clean lint test

clean:				## Go clean and clear all the binaries created.
	@go clean
	@rm -f bin/*

lint:				## Run go linter
	@golangci-lint run -v ./...

test:				## Run all the go tests
	@go test -v -timeout 20s -bench=. ./... -coverprofile coverage.out

verify-makefile:		## Run checkmake to verify Makefile integrity
	@checkmake --debug $(MAKEFILE_LIST)

push-coverage:
	hack/coverage.sh

.PHONY: all clean lint test usage
