#!/usr/bin/make
# Makefile for API microservices
# Copyright (C) 2021 OnLife LLC.

SHELL := /bin/bash
GOPATH := $(shell go env GOPATH)
INCL := -I/usr/local/include

# List of buildable services
SERVICES = auth orders org receipt


default: all


.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get -u github.com/googleapis/gnostic
	go install github.com/googleapis/gnostic/apps/protoc-gen-openapi


# Pass on "all" recipe to all registered services
.PHONY: all
all:
	@(for service in $(SERVICES) ; do \
  		echo "cd $$service" ; \
		$(MAKE) all -C $$service ; \
	done) ;


.PHONY: proto
proto:
	@(for service in $(SERVICES) ; do \
		$(MAKE) proto -C $$service ; \
	done) ;


.PHONY: build
build:
	@(for service in $(SERVICES) ; do \
		$(MAKE) build -C $$service ; \
	done) ;


.PHONY: clean
clean:
	@(for service in $(SERVICES) ; do \
  		echo -n "cd $$service && " ; \
		$(MAKE) clean -C $$service ; \
	done) ;


define generate_openapi =
	$(info Generate OpenAPI files for $(1));
	$(eval exists := $(call dir_exists,$(1)))
endef


define dir_exists =
	$(info Testing OpenAPI directory at $(1))
	$(eval dir := docs/openapi/$(1))
	res := 1;
	if [ -d $(dir) ]; then echo "Dir exists";
		else echo "Dir does not exist"; fi;
	$(info Directory $(dir) result: $(res))
endef



#$(eval res := $(shell (ls -d $(dir) 2>- ; echo $? )))

#ifeq (0,$(shell (ls -d $(dir) 2>- ; echo $?)))
#	$(warning Tmp dir not found);
#	res := 0;
#endif


# Build OpenAPI documentation
.PHONY: openapi
openapi:
	base := "docs/openapi/" ; \
	for service in $(SERVICES) ; do \
		echo $$service ; \
		path :=  $$base ; \
		path += $$service ; \
		echo $$path ; \
		# protoc $(INCLUDE) --openapi_out=openapi --proto_path=. proto/receipt.proto
	done;


.PHONY: docs
docs: openapi
	for service in $(SERVICES) ; do \
		@echo $< $@ $service ; \
		# @todo build recursively
		# @redoc-cli bundle api-receipt.json
	done


test1:
	@-rm -rf tmp.dir
	@-mkdir tmp.dir
	@echo "GOPATH:" $$GOPATH
	find . -name "*.proto" | \
	while read -r file_proto ; do \
		echo $$file_proto ; \
	done
