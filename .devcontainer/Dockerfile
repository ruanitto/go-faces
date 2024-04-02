FROM golang:1.21.6-alpine

RUN apk add --update --no-cache \
    cmake make gcc libc-dev g++ unzip

RUN mkdir /dlib && cd /dlib && \
    wget -c -q  http://dlib.net/files/dlib-19.24.tar.bz2 && \
    tar xf dlib-19.24.tar.bz2

RUN cd /dlib/dlib-19.24 && mkdir build && \
    cd build && cmake .. && \
    cmake --build . --config Release && \
    make install && rm -rf /dlib

# ARG DLIB_BRANCH=v19.24

# RUN wget -c -q https://github.com/davisking/dlib/archive/${DLIB_BRANCH}.tar.gz \
#     && tar xf ${DLIB_BRANCH}.tar.gz \
#     && mv dlib-* dlib \
#     && cd dlib/dlib \
#     && mkdir build \
#     && cd build \
#     && cmake .. \
#     && cmake --build .. --config Release .. \
#     && make \
#     && make install

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

ENV GO111MODULE=off \
    CGO_ENABLED=0