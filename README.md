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

## deploy procedure (current temporary)
```
// boot docker compose
$ make init

// build & debian packaging
$ make dc_build dc_pkg

// deploy
$ scp dst/famiphoto_api_xxxxx.deb 192.168.11.30:/tmp
```

at Server:

```
$ cd /tmp
$ sudo dpkg -i ./famiphoto_api_xxxxxx.deb

// Edit env file
$ sudo nano /etc/famiphoto/env

$ sudo systemctl restart famiphoto_api
$ sudo ssytemctl status famiphoto_api -l 

// if error
$ journalctl _PID=xxxxxx
$ tail -f /var/log/famiphoto/error.log
```