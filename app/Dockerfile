# goバージョン
FROM golang:1.20-alpine

#airのインストール
RUN apk update && apk add git 
RUN go install github.com/cosmtrek/air@latest

# ワーキングディレクトリの変更
WORKDIR /app
# ディレクトリの作成
RUN mkdir -p /go/src/api-server
# ワーキングディレクトリの設定
WORKDIR /go/src/api-server
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/api-server
WORKDIR /go/src/api-server/cmd/api


# コンテナ起動時にアプリケーションを実行
CMD ["go","run","main.go"]

