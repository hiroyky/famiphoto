# Makefile

DST_DIR=./dst
PKG_DIR=$(DST_DIR)/pkg/famiphoto_api
DOCKER=famiphoto
MOCK_TARGETS=$(shell find . -type f -name "*.go" | grep -v "testing/" | grep -v "_test.go" | grep -v "dst/")

init:
	docker compose -f docker-compose.local.yaml build --no-cache
	docker compose -f docker-compose.local.yaml up -d

restart:
	docker compose -f docker-compose.local.yaml stop
	docker compose -f docker-compose.local.yaml up -d

build: build_server build_sub_indexing build_sub_indexing_photos build_register_client

build_prepare:
	go mod tidy
	go mod verify

build_server: build_prepare
	go build -o $(DST_DIR)/app main.go
build_sub_indexing: build_prepare
	go build -o $(DST_DIR)/indexing subsystems/indexing/main.go
build_sub_indexing_photos: build_prepare
	go build -o $(DST_DIR)/indexing_photos subsystems/indexing_photos/main.go
build_register_client: build_prepare
	go build -o $(DST_DIR)/register_client subsystems/register_client/main.go

pkg: build
	mkdir -p $(PKG_DIR)/usr/bin
	cp $(DST_DIR)/app $(PKG_DIR)/usr/bin/famiphoto_api
	cp $(DST_DIR)/indexing_photos $(PKG_DIR)/usr/bin/famiphoto_indexing_photos
	cp $(DST_DIR)/register_client $(PKG_DIR)/usr/bin/famiphoto_register_client
	cp -r pkg/. $(PKG_DIR)
	dpkg-deb --build $(PKG_DIR) $(DST_DIR)

fmt:
	go fmt ./...

test:
	go test ./...

dc_exec:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) bash
dc_exec_indexing_photos:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) ./dst/indexing_photos --env=/go/src/github.com/hiroyky/famiphoto/.env.local --fast=false
dc_exec_indexing:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) ./dst/indexing
dc_fmt:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) make fmt
dc_build:
	docker compose -f docker-compose.local.yaml  exec $(DOCKER) make build
dc_test:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) make test
dc_gen:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) go generate ./...
dc_sqlboil:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) sqlboiler mysql
dc_gengql:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) gqlgen
dc_genmock:
	docker compose -f docker-compose.local.yaml exec $(DOCKER) make mockgen -B -j3
dc_pkg:
	docker compose -f docker-compose.local.yaml exec builder make pkg -B -j3

mockgen: $(MOCK_TARGETS)

$(MOCK_TARGETS):
	mockgen -source $@ -destination testing/mocks/$@

clean:
	rm -rf ./testing/mocks/
	rm -rf ./dst
