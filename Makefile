GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BUF_VERSION=v1.5.0

.PHONY: init
# init env
init:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-saas/kit/cmd/protoc-gen-go-grpc-proxy@ffede0b182a5c92e22da30585ba5a74658aff773
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/go-saas/kit/cmd/protoc-gen-go-errors-i18n/v2@ffede0b182a5c92e22da30585ba5a74658aff773
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.7
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@$(BUF_VERSION)
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@$(BUF_VERSION)



.PHONY: config
# generate internal proto
config:
	buf generate . --path ./private/conf --template buf.config.gen.yaml
	buf generate . --path ./event --template buf.config.gen.yaml

.PHONY: api
# generate api proto
api:
	buf generate . --path ./api

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


.PHONY: all
# generate all
all:
	make init;
	make api;
	make config;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
