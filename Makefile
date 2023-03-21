.PHONY: build run test

export GO111MODULE ?= on

PACKAGES = $(shell go list ./...)
PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY:
db_up:
	@echo "Start MongoDB..."
	@brew services start mongodb-community@4.4

.PHONY:
db_down:
	@echo "Stop MongoDB..."
	@brew services stop mongodb-community@4.4