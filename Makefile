# Makefile

DST_DIR=./dst
DOCKER=famiphoto

init:
	docker compose build --no-cache
	docker compose up -d

restart:
	docker compose stop && docker compose up -d

build:
	go mod tidy
	go mod verify
	go build -o $(DST_DIR)/app main.go

fmt:
	go fmt ./...

test:
	go test ./... -v -count 1

dc_exec:
	docker compose exec $(DOCKER) bash
dc_fmt:
	docker compose exec $(DOCKER) make fmt
dc_build:
	docker compose exec $(DOCKER) make build
dc_test:
	docker compose exec $(DOCKER) make test
dc_gen:
	docker compose exec $(DOCKER) go generate ./...
dc_sqlboil:
	docker compose exec $(DOCKER) sqlboiler mysql
dc_gengql:
	docker compose exec $(DOCKER) gqlgen