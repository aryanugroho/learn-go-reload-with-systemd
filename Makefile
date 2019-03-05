#!/bin/bash

export PKGS=$(shell go list ./... | grep -v vendor/)

build:
	@echo "Building binary..."
	@go build -o go-systemd ./
	@echo "Done."