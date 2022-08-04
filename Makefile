.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

export GO111MODULE=on

OUTPUT ?= .

.PHONY: build
build: build-filter

.PHONY: build-filter
build-filter: build-filter-arm build-filter-amd


.PHONY: build-filter-arm
build-filter-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${OUTPUT}/filter-arm rectangle/cmd/filter

.PHONY: build-filter-amd
build-filter-amd:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT}/filter-amd rectangle/cmd/filter
