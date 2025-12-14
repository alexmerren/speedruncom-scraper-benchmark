GO := go
GOFMT := gofmt

VENDOR_DIR := $(CURDIR)/vendor

.PHONY: help test benchmark vendor clean fmt fetch-data

## help: Print this message
help:
	@fgrep -h '##' $(MAKEFILE_LIST) | fgrep -v fgrep | column -t -s ':' | sed -e 's/## //'

## test: Run tests
test:
	$(GO) test -v

## benchmark: Run benchmark tests
benchmark:
	$(GO) test -bench=.

## vendor: Download the vendored dependencies
vendor:
	$(GO) mod tidy
	$(GO) mod vendor

## clean: Remove any build artifacts
clean:
	rm -rf $(VENDOR_DIR)

## fmt: Format all code for the project
fmt: 
	$(GOFMT) -s -w $(CURDIR)

## fetch-data: Fetch comparison data
fetch-data:
	$(GO) run fetch_benchmark_data.go