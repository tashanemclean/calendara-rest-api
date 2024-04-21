# Makefile 
# VERSION := $(shell git describe --tags --first-parent --abbrev=0)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := genai_rest_api

# Go related variables.
GOHOME := $(shell go env GOPATH)
GO_BASE := $(shell pwd)
GO_BIN := $(GO_BASE)/bin
GO_BINARY := $(GO_BIN)/$(PROJECTNAME)
GO_CLI := $(GO_BASE)/cli

go-build: clean
	@echo "Building binary $(GO_BIN)..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_BINARY) main.go

go-install:
	@echo "Installing go dependencies"
	go mod download

go-binary:
	@echo "Running Go Binary..."
	$(GO_BINARY)

go-upx: go-build
	@echo "Compressing binary with UPX"
	upx --best --lzma $(GO_BINARY)

clean:
	@-rm $(GO_BIN)/$(PROJECTNAME) 2> /dev/null
	go clean

air:
	air

air-docker:
	docker run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(shell pwd):/go/src/$(PACKAGE_NAME) -p $(SERVICE_PORT):80 cosmtrek/air


.PHONY: go-build go-install go-binary air
