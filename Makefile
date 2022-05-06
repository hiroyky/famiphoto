# Makefile

DST_DIR=./dst
DOCKER=famiphoto
MOCK_TARGETS= infrastructures/repositories/adapter.go usecases/adapter.go

init:
	docker compose build --no-cache
	docker compose up -d

restart:
	docker compose stop && docker compose up -d

build: build_server build_sub_import build_sub_indexing

build_prepare:
	go mod tidy
	go mod verify

build_server: build_prepare
	go build -o $(DST_DIR)/app main.go

build_sub_import: build_prepare
	go build -o $(DST_DIR)/import_photos subsystems/import_photos/main.go

build_sub_indexing: build_prepare
	go build -o $(DST_DIR)/indexing subsystems/indexing/main.go

fmt:
	go fmt ./...

test:
	go test ./...

dc_exec:
	docker compose exec $(DOCKER) bash
dc_exec_import:
	docker compose exec $(DOCKER) ./dst/import_photos --base-dir photos/yokoyama/hiro
dc_exec_indexing:
	docker compose exec $(DOCKER) ./dst/indexing
dc_fmt:
	docker compose exec $(DOCKER) make fmt
dc_build:
	docker compose exec $(DOCKER) make build
dc_test:
	docker compose exec $(DOCKER) make test
dc_gen: dc_get_wire
	docker compose exec $(DOCKER) go generate ./...
dc_sqlboil:
	docker compose exec $(DOCKER) sqlboiler mysql
dc_gengql:
	docker compose exec $(DOCKER) gqlgen
dc_genmock:
	docker compose exec $(DOCKER) make mockgen -B -j3

dc_get_wire:
	docker compose exec $(DOCKER) go get github.com/google/wire/cmd/wire@v0.5.0

mockgen: $(MOCK_TARGETS)

$(MOCK_TARGETS):
	mockgen -source $@ -destination testing/mocks/$@
