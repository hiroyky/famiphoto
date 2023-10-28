FROM ubuntu:22.04

WORKDIR /go/src/github.com/hiroyky/famiphoto/
COPY ../.. .
RUN apt-get update
RUN apt install -y software-properties-common
RUN add-apt-repository -y ppa:longsleep/golang-backports
RUN apt-get update
RUN apt install -y devscripts build-essential
RUN apt install -y golang-go
