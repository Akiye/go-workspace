FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && apk add libc-dev gcc make mysql-client

#COPY ./landing-page/back-end/api/go.mod .
#COPY ./landing-page/back-end/api/go.sum .

#RUN go mod download && go mod verify

#COPY . .
#COPY ./back-end/api/entrypoint.sh /entrypoint.sh

RUN cd ./back-end/api && go get && cd -

ENV GOPATH=/app/api/cmd/main:/app/api/pkg/

ENTRYPOINT ["sh", "/entrypoint.sh"]
