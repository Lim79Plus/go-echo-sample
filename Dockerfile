# Docker base Image
FROM golang:1.13.2-alpine
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN go mod download

RUN go build -o app.exe
# ポート
EXPOSE 1323

CMD ./app.exe