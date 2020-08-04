FROM golang:1.14.2-buster

LABEL maintainer="Masafumi Harada"

WORKDIR /root

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y gcc make git build-essential cmake sudo && \
    go get -u -d gocv.io/x/gocv && \
    cd $GOPATH/src/gocv.io/x/gocv && \
    make install && \
    cd /root && \
    git clone https://github.com/haradama/gosenbay && \
    cd /root/gosenbay && \
    go build

WORKDIR /root/gosenbay