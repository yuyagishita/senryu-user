FROM golang:1.14.2-alpine3.11
# ENV sourcesdir /go/src/github.com/microservices-demo/user/
ENV MONGO_HOST mytestdb:27017
ENV HATEAOS user
ENV USER_DATABASE mongodb

RUN mkdir /go/src/work
WORKDIR /go/src/work

ADD . /

RUN apk add --no-cache alpine-sdk
RUN apk update
RUN apk add git

# Golang ホットリロード(freshのインストール)
RUN go get github.com/pilu/fresh
# # Golang 環境構築(任意)
RUN go get  github.com/mdempsky/gocode \
 github.com/uudashr/gopkgs/v2/cmd/gopkgs \
 github.com/ramya-rao-a/go-outline \
 github.com/acroca/go-symbols \
 golang.org/x/tools/cmd/guru \
 golang.org/x/tools/cmd/gorename \
 github.com/cweill/gotests/... \
 github.com/fatih/gomodifytags \
 github.com/josharian/impl \
 github.com/davidrjenni/reftools/cmd/fillstruct \
 github.com/haya14busa/goplay/cmd/goplay \
 github.com/godoctor/godoctor \
 github.com/go-delve/delve/cmd/dlv \
 github.com/stamblerre/gocode \
 github.com/rogpeppe/godef \
 github.com/sqs/goreturns \
 golang.org/x/lint/golint 
