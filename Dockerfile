FROM golang:1.14.2-buster

LABEL maintainer="Masafumi Harada"

ENV OPENCV_VERSION "4.3.0"

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y gcc make git wget build-essential cmake zip && \
    git clone https://github.com/haradama/gosenbay && \
    mkdir -p /tmp/opencv
WORKDIR /tmp/opencv
RUN wget https://github.com/opencv/opencv/archive/${OPENCV_VERSION}.zip && \
    unzip ${OPENCV_VERSION}.zip -d . && \
    mkdir /tmp/opencv/opencv-${OPENCV_VERSION}/build && cd /tmp/opencv/opencv-${OPENCV_VERSION}/build/ && \
    cmake -DOPENCV_GENERATE_PKGCONFIG=ON -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D WITH_FFMPEG=ON -D WITH_TBB=ON .. && \
    make -j "$(nproc)" && \
    make install
WORKDIR /gosenbay
ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:/usr/local/lib
RUN rm -rf /tmp/opencv && \
    go build