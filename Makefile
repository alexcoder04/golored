
NAME = golored

SHELL = /bin/sh
RM = rm
GO = go

OUT_DIR = build

VERSION = $(shell git describe --tags --abbrev=0)
COMMIT = $(shell git rev-parse --short $(VERSION))

GOFLAGS = -ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT_SHA=$(COMMIT)"

# build

build:
	go build -v $(GOFLAGS) -o $(NAME) .

# linux
build-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-amd64 .

build-i386:
	GOOS=linux GOARCH=386 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-386 .

build-arm:
	GOOS=linux GOARCH=arm $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-arm .

build-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-arm64 .

build-all: build-amd64 build-i386 build-arm build-arm64

# install

install:
	$(GO) install $(GOFLAGS) .

# clean

clean:
	$(RM) -rf build
