export GO111MODULE := "on"
project_directory := justfile_directory()
OUTPUT := env_var_or_default("OUTPUT", project_directory)

build: build-filter

build-filter: build-filter-arm build-filter-amd

build-filter-arm:
  CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o {{ OUTPUT }}/filter-arm rectangle/cmd/filter

build-filter-amd:
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o {{ OUTPUT }}/filter-amd rectangle/cmd/filter

gen-pinyin:
  go run {{ project_directory }}/cmd/gen-pinyin {{ project_directory }}/action-name.json {{ project_directory }}/cmd/filter/generated.go
