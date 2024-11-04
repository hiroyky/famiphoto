#!/bin/sh

if [ -e ".env" ]; then
    echo ".env file already exists."
    echo "Do you renew again? (y/n)"
    read RENEW
    if [[ $RENEW != "y" && $RENEW != "Y" ]]; then
      exit 0
    fi
fi

# ask

echo "Enter path to your Photo directory(like: /home/famiphoto/). The path is mounted by docker and famiphoto imports."
read CONTENTS_DIR

# output to .env file

echo "# Environment parameters for famiphoto" > .env
echo "APP_ENV=prod" >> .env
echo "WEB_CLIENT_ID=famiphoto_web" >> .env
echo "WEB_CLIENT_SECRET=`head -c 50 </dev/urandom | base64`" >> .env
echo "PORT=8080" >> .env
echo "" >> .env
echo "## MySQL" >> .env
echo "MYSQL_HOST_NAME=famiphoto_mysqldb" >> .env
echo "MYSQL_PORT=3306" >> .env
echo "MYSQL_DATABASE=famiphoto_db" >> .env
echo "MYSQL_USER=famiphoto" >> .env
echo "MYSQL_PASSWORD=password" >> .env
echo "MYSQL_ROOT_PASSWORD=password" >> .env
echo "" >> .env
echo "## Redis for auth" >> .env
echo "OAUTH_REDIS_HOST_NAME=redis_oauth:6379" >> .env
echo "OAUTH_REDIS_DATABASE=0" >> .env
echo "" >> .env
echo "## Elasticsearch" >> .env
echo "ELASTICSEARCH_ADDRESSES=http://elasticsearch:9200/" >> .env
echo "ELASTICSEARCH_PASSWORD=" >> .env
echo "ELASTICSEARCH_FINGER_PRINT=" >> .env
echo "EXIF_TIMEZONE=Asia/Tokyo" >> .env
echo "" >> .env
echo "## Expire time" >> .env
echo "CC_ACCESS_TOKEN_EXPIRE_IN_SEC=600" >> .env
echo "USER_ACCESS_TOKEN_EXPIRE_IN_SEC=604800" >> .env
echo "LOGIN_TOKEN_EXPIRE_SEC=2592000" >> .env


echo "CONTENTS_DIR=$CONTENTS_DIR" >> .env
echo "HMAC_KEY=`head -c 16 </dev/urandom | base64`" >> .env
echo "ACCESS_TOKEN_HASHED_PREFIX=`head -c 16 </dev/urandom | base64`" >> .env
echo "UPLOAD_TOKEN_HASHED_PREFIX=`head -c 16 </dev/urandom | base64`" >> .env
echo "LOGIN_TOKEN_HASHED_PREFIX=`head -c 16 </dev/urandom | base64`" >> .env

echo "saved .env"