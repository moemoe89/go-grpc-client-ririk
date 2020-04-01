[![Build Status](https://travis-ci.org/moemoe89/practicing-grpc-client-golang.svg?branch=master)](https://travis-ci.org/moemoe89/practicing-grpc-client-golang)
[![codecov](https://codecov.io/gh/moemoe89/practicing-grpc-client-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/practicing-grpc-client-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-grpc-client-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-grpc-client-golang)

# practicing-grpc-client-golang #

Practicing GRPC Client using Golang (Gin Gonic Framework) as Programming Language. This is client implementation from [https://github.com/moemoe89/practicing-grpc-server-golang](https://github.com/moemoe89/practicing-grpc-server-golang) 

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-grpc-client-golang/
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
$ mv <cloned directory> practicing-grpc-client-golang
```

## Running Application with Makefile
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-client-golang
$ make build
```
Run
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-client-golang
$ make run
```
Stop
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-client-golang
$ make stop
```
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Docker Build
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-client-golang
$ make docker-build
```
Docker Up
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-client-golang
$ make docker-up
```
Docker Down
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-client-golang
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