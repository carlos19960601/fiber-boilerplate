PWD=$(shell pwd)

wire:
	wire gen $(PWD)/cmd/server/wire

build: wire
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

run: build
	bin/server

wire-migration:
	wire gen $(PWD)/cmd/migration/wire

build-migration: wire-migration
	go build -ldflags="-s -w" -o ./bin/server ./cmd/migration

run-migration: build-migration  
	bin/server
