# EasyCreateRotenAppWeb
4Iで開発した露店の回転を効率化するアプリです！！

# アーキテクチャ図
![aws](https://github.com/user-attachments/assets/6cc7a495-e9a5-4cca-b92e-9d6d9f0d16fd)

# 使用技術
## フロントエンド
- フレームワーク: Next.js (App Router)
- CSS: Tailwind CSS
- パッケージマネージャー: npm
- 状態管理: Jotai
- API通信: Next.jsのfetch
## バックエンド
- フレームワーク: Golang (Echo)
- データベース: MySQL
- ORM: Gorm
- API: REST API
- 認証: JWT
- デプロイ: AWS (ECR)
- CDN: AWS CloudFront
- コンテナ: Docker

## Gitの運用ルール
### Issue
機能追加やバグ、UI 等の改善時にIssueを立ててから作業する
## ブランチ名
### Issue番号を含める
feature（新機能開発）<br>
例：feature/#8_add_menu <br>
fix（バグ修正）<br>
例：fix/#21_fix_bug <br>

## コミットメッセージ
### Prefix（接頭辞）をつける<br>
覚えることがめんどくさいので以下3つを使い分ける<br>
feat(新規追加)<br>
fix(修正）<br>
chore(動作に影響しないファイルの編集)<br>

### Issue 番号を末尾につける
コミットメッセージの末尾に固有IDを含めて Issueと関連づける
例：update: ◯◯ なため、△△ を追加する #3

## Development Environment
開発環境をdocker-composeで構築しています。

### Requirements
* Go
* docker / docker-compose

### Start
```
docker compose up -d
```

### Shutdown
```
docker compose down
```


