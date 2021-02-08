.PHONY: build

build:
		go build -v ./cmd/avtask
.DEFAULT_GOAL := build
