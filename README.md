# Famiphoto
Let's manage my photos at home server.

## What is this?

**Famiphoto**

A Server application for managing photos at home server.
You know many people manage photos on cloud storage services such as Google Photos, Amazon Photos, OneDrive or Dropbox.
But, do you have felt that I want to manage my home server? If on my home server, subscription fee is nothing, free, and my dream.
I also think so. So I am developing this application.

## GraphQL API

This repository is GraphQL API component. and Frontend web system is [Famiphoto frontend](https://github.com/hiroyky/famiphoto_frontend).

## I'm sorry.

I'm developing this app as my hobby. so Please wait a little(??).

## procedure to deploy as Docker container (recommended)
### Before install
Please set up docker & docker compose on your server machine.

### Setup Famiphoto API
1. Execute `setup_env.sh` to configure .env file. 
Then, specify the path to your Photo directory which you want Famiphoto to mount. 
```:sh
./setup_env.sh
```
2. docker compose begin
```sh:
docker compose build
docker compose up -d
```

3. Check API started.
```
curl http://localhost:8080/status.html
```

### Setup Famiphoto Frontend
1. Git clone famiphoto_frontend
```
git clone git@github.com:hiroyky/famiphoto_frontend.git
cd famiphoto_frontend
```
2. Execute `setup_env.sh` to configure .env file.
```
./setup_env.sh
```

3. docker compose begin
```sh:
docker compose build
docker compose up -d
```

4. Access to Famiphoto frontend