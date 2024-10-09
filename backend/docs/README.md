### 参考にした記事

https://zenn.dev/yuki_tu/articles/bcb926b0ebae47

### 環境の選択

```shell
# deploy環境でビルド
$ docker build --no-cache --target deploy ./

# dev環境でビルド
$ docker build --no-cache --target dev ./

```

### Docker 起動手順

1. Docker を立ち上げる
2. docker compose up -d --build
