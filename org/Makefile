#!/usr/bin/make
# Makefile for Org microservice
# Copyright (C) 2021 OnLife LLC.

SHELL := /bin/bash
GOPATH := $(shell go env GOPATH)
INCLUDE := -I/usr/local/include

default: build

.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro

.PHONY: proto
proto:
	protoc $(INCLUDE) --proto_path=. --micro_out=. --go_out=:. proto/org.proto

.PHONY: build
build: org

org:
	go build -o org *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t org:latest

.PHONY: clean
clean:
	rm -f ./org

# Clean and build from scratch
.PHONY: all
all: clean proto build test
