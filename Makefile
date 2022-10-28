GO = $(shell which go)

.PHONY: ensurebin build build-linux clean lint test fmt

CMD_PATH = ./cmd/slackthing
BIN_OUT = ./bin/slackthing

build : ensurebin
	$(GO) build -o $(BIN_OUT) $(CMD_PATH)

build-linux : ensurebin
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN_OUT)-linux $(CMD_PATH)

ensurebin :
	if ! test -d ./bin; then mkdir -p ./bin; fi

clean :
	@echo "not implemented"

lint :
	@golangci-lint run --config=.golangci.yml

lint-fix :
	@golangci-lint run --config=.golangci.yml --fix

test :
	@echo "not implemented"

fmt :
	gofmt -s -w ./