### Makefile for version

GOPATH ?= $(shell go env GOPATH)

ifeq "$(GOPATH)" ""
	$(error Please set the environment variable GOPATH before running `make`)
endif

GO:=go
GOBUILD   := GOPATH=$(GOPATH) $(GO) build
	LDFLAGS += -X "version_demo/internal/util.Version=$(shell git rev-parse HEAD)"
	LDFLAGS += -X "version_demo/internal/util.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
	LDFLAGS += -X "version_demo/internal/util.Branch=$(shell git rev-parse --abbrev-ref HEAD)"
all:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o version main.go
clean:
	rm version
