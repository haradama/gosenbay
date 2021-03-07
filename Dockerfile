FROM gocv/opencv:4.5.1

LABEL maintainer="Masafumi Harada"

WORKDIR /root

RUN apt update -y && \
    git clone https://github.com/haradama/gosenbay && \
    cd /root/gosenbay && \
    go install

WORKDIR /