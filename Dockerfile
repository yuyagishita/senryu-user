FROM golang:1.14.2-alpine3.11

RUN mkdir /go/src/work
WORKDIR /go/src/work

ADD . /

RUN apk add --no-cache alpine-sdk
RUN apk update
RUN apk add git

# Golang ホットリロード(freshのインストール)
RUN go get github.com/pilu/fresh
# Golang 環境構築(任意)
RUN go get github.com/go-delve/delve/cmd/dlv \
    github.com/rogpeppe/godef \ 
    golang.org/x/tools/cmd/goimports \
    golang.org/x/tools/cmd/gorename \
    sourcegraph.com/sqs/goreturns \
    github.com/ramya-rao-a/go-outline \
    golang.org/x/tools/gopls@latest
