FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-grpc-client-ririk

WORKDIR /go/src/github.com/moemoe89/go-grpc-client-ririk

COPY . /go/src/github.com/moemoe89/go-grpc-client-ririk

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-grpc-client-ririk
