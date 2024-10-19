# ベースイメージ
FROM golang:1.23-alpine AS development

# 作業ディレクトリの設定
WORKDIR /app

# air のインストール
RUN go install github.com/cosmtrek/air@v1.49.0

# Goモジュールファイルをコピー
COPY ../go.mod ../go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY ../ .

# air 設定ファイルのコピー
COPY .air.toml .

# air を起動コマンドとして設定
CMD ["air"]
