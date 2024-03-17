# 名前
golang-clean-architecuture

## 概要
go　x　Ginを使った簡単なCRUD webAPIです。

## フォルダ構成
```
.
├── README.md
└── app
    ├── cmd
    │   └── main.go
    ├── docker
    │   ├── Dockerfile
    │   ├── docker-compose.yml
    │   └── tmp
    │       └── build-errors.log
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── domain
    │   │   ├── entity
    │   │   │   └── post.go
    │   │   └── repository
    │   │       └── post_repository.go
    │   ├── infrastructure
    │   │   ├── database
    │   │   │   └── database.go
    │   │   ├── repository
    │   │   │   ├── post_repository.go
    │   │   │   └── post_repository_test.go
    │   │   └── router.go
    │   ├── interfaces
    │   │   ├── dto
    │   │   │   └── post_dto.go
    │   │   └── handler
    │   │       └── post_handler.go
    │   └── usecase
    │       └── post_usecase.go
    ├── migrations
    │   ├── 20240310151506-posts.sql
    │   └── dbconfig.yml
    ├── pkg
    │   ├── const.go
    │   ├── errors.go
    │   ├── middleware.go
    │   ├── test_utils.go
    │   ├── utils.go
    │   └── validation.go
    ├── post_cmd.txt
    ├── pre_cmd.txt
    └── tmp
        └── main
```

## 動作環境
- macOS(intel Sonoma)
- go-version 1.20
- Docker-version 24.0.7
  
## 使用方法
/app/docker 配下でdocker compose up するだけで使用できます。

## エンドポイント
localhost:8080
- 登録: /add_post
- 更新: /edit_post/:id
- 削除: /delete_post/:id
- 一覧取得: /get_posts

## 特徴
ミドルウェアでのログ設定やバリデーション処理を実装しています。
