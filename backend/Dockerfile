# 開発環境(dev)
FROM golang:1.22-alpine AS dev
WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends curl=7.68.0-1ubuntu2.6 \
    && go install github.com/cosmtrek/air@v1.49.0

# go.modとgo.sumを個別にコピーしてキャッシュを有効化
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# dev環境ではホットリロードを使って実行
CMD ["air"]

# ビルド環境(build)
FROM golang:1.22-alpine AS build
WORKDIR /app

# go.modとgo.sumを個別にコピーしてキャッシュを効率化
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main

# デプロイ環境(deploy)
FROM alpine:3.19 AS deploy
WORKDIR /root/

# ビルド済みバイナリをコピー
COPY --from=build /app/main .

CMD ["./main"]