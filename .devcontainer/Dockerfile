FROM qmcgaw/godevcontainer:alpine

RUN apk add --update --no-cache \
    cmake make gcc libc-dev g++ unzip

RUN mkdir /dlib && cd /dlib && \
    wget -c -q  http://dlib.net/files/dlib-19.24.tar.bz2 && \
    tar xf dlib-19.24.tar.bz2

RUN cd /dlib/dlib-19.24 && mkdir build && \
    cd build && cmake .. && \
    cmake --build . --config Release && \
    make install && rm -rf /dlib