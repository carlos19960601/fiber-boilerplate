PWD=$(shell pwd)

wire:
	wire gen $(PWD)/cmd/server/wire

build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server