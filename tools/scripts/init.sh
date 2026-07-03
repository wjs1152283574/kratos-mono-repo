#!/bin/bash
# 环境初始化脚本
# 安装 protoc 相关工具和依赖

set -e

echo "==> Installing protoc tools..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install github.com/envoyproxy/protoc-gen-validate@latest
go install github.com/google/wire/cmd/wire@latest
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

echo "==> All tools installed!"
echo ""
echo "Installed tools:"
echo "  - protoc-gen-go"
echo "  - protoc-gen-go-grpc"
echo "  - protoc-gen-go-http"
echo "  - protoc-gen-go-errors"
echo "  - protoc-gen-openapiv2"
echo "  - protoc-gen-validate"
echo "  - wire"
echo "  - kratos"
