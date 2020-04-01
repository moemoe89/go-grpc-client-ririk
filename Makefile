.PHONY: build run swag goose stop

test:
	@go test -v -cover -covermode=atomic ./...

build:
	@go build

run:
	@cp config-local.json config.json
	@go build -o app
	@nohup ./app &

stop:
	@kill -9 `lsof -t -i:8794`

docker-build:
	@cp -rf config-docker.json config.json
	@docker-compose build

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down
