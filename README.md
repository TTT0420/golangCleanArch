# 名前
golang-clean-architecuture

## 概要
go　x　Ginを使った簡単なCRUD webAPIです。

## 動作環境
- macOS(intel Sonoma)
- go-version 1.20
- Docker-version 24.0.7
  
## Usage
/app/docker 配下でdocker compose up するだけで使用できます。

エンドポイントは
- 登録: /add_post
- 更新: /edit_post/:id
- 削除: /delete_post/:id
- 一覧取得: /get_posts

## Features
ミドルウェアでのログ設定やバリデーション処理を実装しています。
