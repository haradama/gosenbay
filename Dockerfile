FROM ubuntu:18.04

LABEL maintainer="Masafumi Harada"

RUN apt-get update && \
    apt-get upgrade -y && \
