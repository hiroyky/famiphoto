# Makefile

DST_DIR=./dst
PKG_DIR=$(DST_DIR)/pkg/famiphoto_api
DOCKER=famiphoto
MOCK_TARGETS=$(shell find . -type f -name "*.go" | grep -v "testing/" | grep -v "_test.go" | grep -v "dst/")

init:
	docker compose build --no-cache
	docker compose up -d

restart:
	docker compose stop && docker compose up -d

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
	docker compose exec $(DOCKER) bash
dc_exec_indexing_photos:
	docker compose exec $(DOCKER) ./dst/indexing_photos
dc_exec_indexing:
	docker compose exec $(DOCKER) ./dst/indexing
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
dc_genmock:
	docker compose exec $(DOCKER) make mockgen -B -j3
dc_pkg:
	docker compose exec builder make pkg -B -j3

mockgen: $(MOCK_TARGETS)

$(MOCK_TARGETS):
	mockgen -source $@ -destination testing/mocks/$@

clean:
	rm -rf ./testing/mocks/
	rm -rf ./dst

prodinit:
	docker-compose --file docker-compose-prod.yaml build --no-cache
	docker-compose --file docker-compose-prod.yaml up -d