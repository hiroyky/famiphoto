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

echo "Enter base URL of Famiphoto frontend. (like: https://famiphoto.myhome.com)"
read BASE_URL

WEB_CLIENT_SECRET=`docker exec famiphoto env | grep WEB_CLIENT_SECRET | sed "s/WEB_CLIENT_SECRET=//"`
if [[ $WEB_CLIENT_SECRET == "" ]];then
  echo "Failed to get WEB_CLIENT_SECRET from Famiphoto API docker container."
  echo "Enter famiphoto api's env WEB_CLIENT_SECRET value:"
  read WEB_CLIENT_SECRET
fi

# output .env
echo "#Famihoto frontend env" > .env
echo "NUXT_IS_DEBUG=false" >> .env
echo "NUXT_API_BASE_URL=http://host.docker.internal:8080" >> .env
echo "NUXT_CLIENT_SECRET=$WEB_CLIENT_SECRET" >> .env
echo "NUXT_PUBLIC_BASE_URL=$BASE_URL" >> .env
echo "NUXT_SESSION_SECRET=`head -c 50 </dev/urandom | base64`" >> .env
