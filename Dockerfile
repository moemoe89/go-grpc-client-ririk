FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-grpc-client-golang

WORKDIR /go/src/github.com/moemoe89/practicing-grpc-client-golang

COPY . /go/src/github.com/moemoe89/practicing-grpc-client-golang

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-grpc-client-golang
