FROM golang:1.18.4-alpine3.16

RUN go install github.com/cosmtrek/air@v1.40.4

WORKDIR /opt/app/api

CMD ["air"]
