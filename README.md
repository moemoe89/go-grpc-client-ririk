[![Build Status](https://travis-ci.org/moemoe89/go-grpc-client-ririk.svg?branch=master)](https://travis-ci.org/moemoe89/go-grpc-client-ririk)
[![codecov](https://codecov.io/gh/moemoe89/go-grpc-client-ririk/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/go-grpc-client-ririk)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/go-grpc-client-ririk)](https://goreportcard.com/report/github.com/moemoe89/go-grpc-client-ririk)

# GO-GRPC-CLIENT-RIRIK #

Practicing GRPC Client using Golang as Programming Language. This is client implementation from [https://github.com/moemoe89/go-grpc-server-tisa](https://github.com/moemoe89/go-grpc-server-tisa) 

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ go-grpc-client-ririk/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> go-grpc-client-ririk
```

## Running Application with Makefile
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ cd $GOPATH/src/github.com/moemoe89/go-grpc-client-ririk
$ make build
```
Run
```
$ cd $GOPATH/src/github.com/moemoe89/go-grpc-client-ririk
$ make run
```
Stop
```
$ cd $GOPATH/src/github.com/moemoe89/go-grpc-client-ririk
$ make stop
```
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Docker Build
```
$ cd $GOPATH/src/github.com/moemoe89/go-grpc-client-ririk
$ make docker-build
```
Docker Up
```
$ cd $GOPATH/src/github.com/moemoe89/go-grpc-client-ririk
$ make docker-up
```
Docker Down
```
$ cd $GOPATH/src/github.com/moemoe89/go-grpc-client-ririk
$ make docker-down
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## Reference

Thanks to this medium [link](https://toolbox.kurio.co.id/implementing-grpc-service-in-golang-afb9e05c0064) for sharing the great article

## License

MIT