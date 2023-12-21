#!/usr/bin/make -f

BUILD_DIR ?= $(CURDIR)/build
PACKAGE_CMD_DIR := $(CURDIR)/cmd

.PHONY: build install

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

install: go.sum
	go install -mod=readonly .

build:
	go build -mod=readonly -o $(BUILD_DIR)/hypersign-load-test .