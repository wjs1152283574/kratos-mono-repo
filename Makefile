GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find app -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
name=demo
app=demo

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get -u github.com/envoyproxy/protoc-gen-validate

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
               --validate_out=paths=source_relative,lang=go:. \
               --openapiv2_out . \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: builds
# generate builds
builds:
	cd app/${app}/service && make build

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;

.PHONY: app
# newapp
app:
	kratos proto add api/$(name)/service/v1/$(name).proto && \
	kratos proto client api/$(name)/service/v1/$(name).proto && \
	cd app && mkdir $(name) && cd ./$(name) && mkdir service && cd ./service && mkdir cmd && mkdir ./cmd/server && \
	cd ./cmd/server touch main.go && echo "package main" >> ./main.go && cd .. && cd .. \
	mkdir configs && mkdir internal && cd internal && mkdir biz && \
	cd biz && touch biz.go && echo "package biz" >> ./biz.go && cd .. && \
	mkdir conf && mkdir data && \
	cd data && touch data.go && echo "package data" >> ./data.go && cd .. && \
	mkdir server && \
	cd server && touch server.go && echo "package server" >> ./server.go && cd .. && \
	mkdir service && cd service && \
	touch service.go && echo "package service" >> ./service.go && cd .. && \
	cd .. && touch .gitignore && touch generate.go && echo "package generate" >> ./generate.go && \
	touch Makefile && echo "include ../../../app_makefile" >> ./Makefile && touch README.MD && cd ../../../ && \
	kratos proto server api/$(name)/service/v1/$(name).proto -t app/$(name)/service/internal/service

.PHONY: initdb
# initdb
initdb:
	cd deploy/docker-compose && docker-compose up -d

.PHONY: run
# run
run:
	cd app/${app}/service && make run

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
