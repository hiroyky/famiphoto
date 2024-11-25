# Makefile

init:
	docker compose -f docker-compose.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.yaml --env-file ./.env up  -d

restart:
	docker compose -f docker-compose.yaml stop
	docker compose -f docker-compose.yaml up -d


local_init:
	docker compose -f docker-compose.local.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.local.yaml --env-file ./.env up  -d

local_restart:
	docker compose -f docker-compose.local.yaml stop
	docker compose -f docker-compose.local.yaml up -d

dc_exec_api:
	docker compose -f docker-compose.local.yaml exec famiphoto_api bash
