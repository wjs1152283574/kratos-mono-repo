#!/bin/bash
# Proto 代码生成脚本
# 用法: bash tools/scripts/proto.sh [service_name]
# 示例: bash tools/scripts/proto.sh          # 生成所有
#       bash tools/scripts/proto.sh shop      # 只生成 shop

set -e

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)
cd "$ROOT_DIR"

if [ -z "$1" ]; then
    echo "==> Generating all API proto files..."
    API_PROTO_FILES=$(find api -name "*.proto")
else
    echo "==> Generating API proto files for: $1"
    API_PROTO_FILES=$(find "api/$1" -name "*.proto" 2>/dev/null || true)
    if [ -z "$API_PROTO_FILES" ]; then
        echo "Error: No proto files found for $1"
        exit 1
    fi
fi

# 生成 API 代码到 gen/
protoc --proto_path=. \
       --proto_path=./third_party \
       --go_out=paths=source_relative:. \
       --go-http_out=paths=source_relative:. \
       --go-grpc_out=paths=source_relative:. \
       --go-errors_out=paths=source_relative:. \
       --validate_out=paths=source_relative,lang=go:. \
       --openapiv2_out . \
       $API_PROTO_FILES

echo "==> API proto generation done!"

# 生成内部配置 proto
echo "==> Generating internal config proto files..."
if [ -z "$1" ]; then
    INTERNAL_PROTO_FILES=$(find app -name "*.proto")
else
    INTERNAL_PROTO_FILES=$(find "app/$1" -name "*.proto" 2>/dev/null || true)
fi

if [ -n "$INTERNAL_PROTO_FILES" ]; then
    protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           $INTERNAL_PROTO_FILES
    echo "==> Internal config proto generation done!"
else
    echo "==> No internal proto files found, skipping."
fi

echo "==> All proto generation complete!"
