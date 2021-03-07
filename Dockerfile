FROM gocv/opencv:4.5.1

LABEL maintainer="Masafumi Harada"

WORKDIR /tmp

RUN apt update -y && \
    git clone https://github.com/haradama/gosenbay && \
    cd /tmp/gosenbay && \
    go install && \
    rm -rf /tmp/gosenbay

WORKDIR /