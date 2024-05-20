# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOTEST=$(GOCMD) test
BINARY_NAME=gerrit-translator-cdevents
PKG_DIR=./pkg/

all: fmt vet test build

build:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -o $(BINARY_NAME) $(PKG_DIR)

test:
	$(GOTEST) -v ./...

fmt:
	$(GOFMT) ./...

vet:
	$(GOVET) ./...

clean:
	rm -f $(BINARY_NAME)
