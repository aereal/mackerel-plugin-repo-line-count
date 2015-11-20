VERSION = $$(git describe --tags --always --dirty)
BUILD_LDFLAGS = "-X main.Version=$(VERSION)"
BIN = mackerel-plugin-repo-line-count
BUILD_DIR = ./build

all: build

deps:
	go get -d -v -t ./...

build: deps
	mkdir -p $(BUILD_DIR)
	go build -ldflags=$(BUILD_LDFLAGS) -o $(BUILD_DIR)/$(BIN)

fmt:
	gofmt -w .

test: deps
	go test ./...
