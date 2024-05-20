FROM golang:1.21 as local

ENV GO111MODULE=on
WORKDIR /go/src/github.com/hiroyky/famiphoto
COPY ../.. .
RUN apt-get update
RUN apt-get install -y libmagickwand-dev
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
RUN go install github.com/99designs/gqlgen@v0.17.38
RUN go get github.com/volatiletech/sqlboiler/boil
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN mkdir -p /var/famiphoto
RUN mkdir -p /tmp
RUN mkdir -p /var/www/famiphoto
