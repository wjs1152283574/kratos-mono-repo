FROM golang:1.22 AS builder

ARG APP_NAME

COPY . /src
WORKDIR /src

RUN go mod download

WORKDIR /src/app/${APP_NAME}

RUN CGO_ENABLED=0 GOPROXY=https://goproxy.cn go build -ldflags "-X main.Version=$(git describe --tags --always)" -o ./bin/server ./cmd/server/

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_NAME}/bin/server /app/server

WORKDIR /app

EXPOSE 8000 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
