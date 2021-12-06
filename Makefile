OUTPUT = "appli.ed"
ALL_SOURCE = $(shell find . -name "*.go")
PROD_SOURCE = $(shell find . -name "*.go" -type f ! -path "*_test.go")

all: test build

build: out/$(OUTPUT)

lint:
	@if [[ -z $(shell which golangci-lint) ]]; then echo "Install golangci-lint" >&2; exit 1; fi
	golangci-lint run

out/$(OUTPUT): $(PROD_SOURCE)
	go build -o out/$(OUTPUT) ./main.go

test: $(ALL_SOURCE)
	go run github.com/onsi/ginkgo/ginkgo ./...

.PHONY: all build clean lint test
