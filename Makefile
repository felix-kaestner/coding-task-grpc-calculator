#!make

PKG:=github.com/felix-kaestner/calculator

build: build-client build-server

build-client:
	@go build $(PKG)/cmd/client

build-server:
	@go build $(PKG)/cmd/server

fmt:
	@gofmt -w .

test:
	@go test -v ./... -race

coverage:
	@go test -v ./... -race -cover -coverprofile=coverage.out -covermode=atomic
	@go tool cover -html=coverage.out -o coverage.html

clean: 
	@rm client server coverage.out coverage.html