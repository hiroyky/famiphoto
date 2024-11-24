# Makefile

init:
	docker compose -f docker-compose.local.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.local.yaml --env-file ./.env up  -d

restart:
	docker compose -f docker-compose.local.yaml stop
	docker compose -f docker-compose.local.yaml up -d
