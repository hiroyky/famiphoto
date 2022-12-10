#!/bin/sh

## generate .env for famiphoto_web

mkdir -p ./dst
WEB_ENV=./dst/.env.web.prod

if [ ! -e ${WEB_ENV} ]; then
  touch ${WEB_ENV}
  echo "API_BASE_URL=http://famiphoto:8080" >> ${WEB_ENV}
  echo "IS_DEBUG=0" >> ${WEB_ENV}

  CLIENT_SECRET=`cat /dev/urandom |LC_CTYPE=C tr -dc '[:alnum:]' | fold -w 50 |head -n 1`
  SESSION_SECRET=`cat /dev/urandom |LC_CTYPE=C tr -dc '[:alnum:]' | fold -w 32 |head -n 1`

  echo "CLIENT_SECRET=${CLIENT_SECRET}" >> ${WEB_ENV}
  echo "SESSION_SECRET=${SESSION_SECRET}" >> ${WEB_ENV}
fi

## begin docker compose

docker compose --file docker-compose-prod.yaml build --no-cache
docker compose --file docker-compose-prod.yaml up -d

if [ -n "${CLIENT_SECRET}" ]; then
  docker compose exec famiphoto ./dst/register_client --client-id famiphoto --name "FAMIPHOTO" --client-secret ${CLIENT_SECRET}
fi
