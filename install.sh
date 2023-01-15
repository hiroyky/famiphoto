#!/bin/sh

function random_str() {
    cat /dev/urandom |LC_CTYPE=C tr -dc '[:alnum:]' | fold -w $1 |head -n 1
}

CLIENT_SECRET=`random_str 50`

mkdir -p ./dst

## generate .env for famiphoto_web

WEB_ENV=./dst/.env.web.prod

echo "API_BASE_URL=http://famiphoto:8080" > ${WEB_ENV}
echo "IS_DEBUG=0" >> ${WEB_ENV}
echo "CLIENT_ID=famiphoto_web" >> ${WEB_ENV}
echo "CLIENT_SECRET=${CLIENT_SECRET}" >> ${WEB_ENV}
echo "SESSION_SECRET=`random_str 32`" >> ${WEB_ENV}

## generate .env for famiphoto api
API_ENV=./dst/.env.prod

echo "# APP ENV Prod Appendix" > ${API_ENV}
echo "WEB_CLIENT_SECRET=${CLIENT_SECRET}" >> ${API_ENV}
echo "MYSQL_PASSWORD=`random_str 32`" >> ${API_ENV}
echo "HMAC_KEY=`random_str 32`" >> ${API_ENV}
echo "ACCESS_TOKEN_HASHED_PREFIX=`random_str 16`" >> ${API_ENV}
echo "UPLOAD_TOKEN_HASHED_PREFIX=`random_str 16`" >> ${API_ENV}
echo "LOGIN_TOKEN_HASHED_PREFIX=`random_str 16`" >> ${API_ENV}

## begin docker compose

docker compose --file docker-compose-prod.yaml build --no-cache
docker compose --file docker-compose-prod.yaml up -d

docker compose exec famiphoto elasticsearch/create_index.sh
