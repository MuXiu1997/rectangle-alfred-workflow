.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

export GO111MODULE=on

OUTPUT ?= .

build: build-action build-filter

build-action:
	go build -o $(OUTPUT)/action rectangle/cmd/action

build-filter:
	go build -o $(OUTPUT)/filter rectangle/cmd/filter
