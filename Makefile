#!/usr/bin/make -f

# Disable all default suffixes
.SUFFIXES:

# ----- Aliases
.PHONY: default

default: todo


# ----- Protobuf
proto_module_prefix := $(shell go list -m -tags tools)
proto_sources := $(shell find proto/src -type f -name '*.proto')
proto_targets := $(proto_sources:proto/src/%.proto=proto/gen/%.pb.go)

.PHONY: proto proto.clean

proto/gen/%.pb.go: proto/src/%.proto
	$(info Compiling protobuf '$@')
	@mkdir -p $(@D)
	@protoc \
		--proto_path=proto/src \
		--go_out=. --go_opt=module=$(proto_module_prefix) \
		$<

proto: $(proto_targets) ## Compile protobufs

proto.clean: ## Clean protobuf artifacts
	$(info Cleaning protobuf artifacts)
	@[ -d proto/gen ] && rm -r proto/gen/* 2> /dev/null || true


# ----- Tooling
.PHONY: fmt lint tidy

fmt: ## Format
	$(info Formatting)
	@go fmt ./...

lint: ## Lint
	$(info Linting)
	@golint ./...

tidy: ## Tidy go modules
	$(info Tidying)
	@go mod tidy -compat=1.17


# ----- Help
.PHONY: help

help: ## Show help information
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST);

print-%: ; @echo "$($*)"