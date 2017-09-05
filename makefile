### Makefile for version

GOPATH ?= $(shell go env GOPATH)

ifeq "$(GOPATH)" ""
	$(error Please set the environment variable GOPATH before running `make`)
endif

GO:=go
GOBUILD   := GOPATH=$(GOPATH) $(GO) build
	# why don't need full package path, e.g. $prejct/packagename
	# LDFLAGS += -X "version_demo.main.Version=$(shell git rev-parse HEAD)"
	# LDFLAGS += -X "version_demo.main.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
	# LDFLAGS += -X "version_demo.main.Branch=$(shell git rev-parse --abbrev-ref HEAD)"

	LDFLAGS += -X "main.Version=$(shell git rev-parse HEAD)"
	LDFLAGS += -X "main.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
	LDFLAGS += -X "main.Branch=$(shell git rev-parse --abbrev-ref HEAD)"
all:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o version main.go
clean:
	rm version
