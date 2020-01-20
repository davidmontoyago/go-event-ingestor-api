# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=on
GOOS?=darwin
GOARCH=amd64
MAX_QUEUE=100000
MAX_WORKERS=10000

all: test build

build:
	go mod vendor
	$(GOBUILD) ./

test:
	$(GOTEST) ./

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./pkg/...
	$(GOCMD) fmt ./main.go

run:
	rm ./go-event-ingestor-api
	make build
	MAX_QUEUE=$(MAX_QUEUE) MAX_WORKERS=$(MAX_WORKERS) ./go-event-ingestor-api

pre-reqs:
	brew install hey
