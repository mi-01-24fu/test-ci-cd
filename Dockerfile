FROM golang:1.21.1-alpine
# 作業ディレクトリの設定
WORKDIR /app

# ソースコードのコピー
COPY . .

# プログラムのビルド
RUN go build -o main .

# コンテナ起動時に実行されるコマンド
CMD ["/app/main"]

# コンテナがリッスンするポートの指定
EXPOSE 8080