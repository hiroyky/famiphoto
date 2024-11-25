# Makefile

DST_DIR=./dst
MOCK_TARGETS=$(shell find . -type f -name "*.go" | grep -v "testing/" | grep -v "_test.go" | grep -v "dst/")

build: build_server build_indexing

build_prepare:
	go mod tidy
	go mod verify

build_server: build_prepare
	go build -o $(DST_DIR)/api apps/api/main.go
build_indexing: build_prepare
	go build -o $(DST_DIR)/indexing apps/indexing/main.go

fmt:
	go fmt ./...

test:
	go test ./...

gen: sqlboiler gengql mockgen
sqlboiler:
	sqlboiler mysql
gengql:
	gqlgen

mockgen: $(MOCK_TARGETS)

$(MOCK_TARGETS):
	mockgen -source $@ -destination testing/mocks/$@

clean:
	rm -rf ./testing/mocks/
	rm -rf ./dst
