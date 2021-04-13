#!/usr/bin/make
# Makefile for Auth microservice
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
	protoc $(INCLUDE) --proto_path=. --micro_out=. --go_out=:. proto/auth.proto

.PHONY: build
build: auth

auth:
	go build -o auth *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t auth:latest

.PHONY: clean
clean:
	rm -f ./auth

# Clean and build from scratch
.PHONY: all
all: clean proto build test
