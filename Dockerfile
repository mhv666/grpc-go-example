FROM golang:1.17.2-bullseye

WORKDIR /go/src/github.com/user/myProject/app/main

# RUN apt-get update && \
#     apt-get -y install git unzip build-essential autoconf libtool
# RUN git clone https://github.com/google/protobuf.git && \
#     cd protobuf && \
#     ./autogen.sh && \
#     ./configure && \
#     make && \
#     make install && \
#     ldconfig && \
#     make clean && \
#     cd .. && \
#     rm -r protobuf

# RUN go get google.golang.org/protobuf
# RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

COPY ./app /go/src/github.com/user/myProject/app

COPY go.mod ./
COPY go.sum ./

WORKDIR /go/src/github.com/user/myProject/app/main/proto
RUN protoc  --proto_path=./ --go_out=./ *.proto

WORKDIR /go/src/github.com/user/myProject/app/main

RUN go build 