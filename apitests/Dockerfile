FROM golang:1.13.1-alpine

RUN apk update && apk add musl dep git

WORKDIR $GOPATH/src/golang-playground/apitests
COPY . .

RUN ["echo", "$GOPATH"]

RUN dep ensure -v

RUN go build -o $GOPATH/bin/app

EXPOSE 8081

ENTRYPOINT [ "entrypoint.sh" ]
