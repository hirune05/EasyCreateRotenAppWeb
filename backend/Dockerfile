# マルチステージビルド利用によるDockerfileの構築

# ビルド用環境
# ----------------------------------------------
FROM --platform=linux/x86_64 golang:1.23.2-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# 本番環境
# ----------------------------------------------
FROM --platform=linux/x86_64 debian:bullseye-slim AS deploy

# X509: Certificate Signed by Unknown Authorityエラーを回避する
RUN apt-get update \
    && apt-get install -y --force-yes --no-install-recommends apt-transport-https curl ca-certificates \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

COPY --from=deploy-builder /app/app .

CMD ["./app"]


# ローカルでの開発環境
# ----------------------------------------------
FROM --platform=linux/x86_64 golang:1.23.2-alpine3.20 AS dev

WORKDIR /app

# 軽量にするためイメージはalpineを利用
RUN apk update && apk add alpine-sdk && apk add jq
# air ライブリロード（変更のたびにサーバーを起動する必要がなくなる）を行ってくれるモジュール
RUN go install github.com/air-verse/air@latest
# sqldef マイグレーションツール（新規作成DDLを判別し差分適用DDL文を自動で生成）
# RUN go install github.com/k0kubun/sqldef/cmd/mysqldef@latest

CMD ["go", "run", "main.go"]