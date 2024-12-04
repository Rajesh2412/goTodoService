FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN  CGO_ENABLED=0 go build -o goTodoService ./cmd

RUN chmod +x /app/goTodoService

FROM alpine:latest 

RUN mkdir /app

WORKDIR /app

COPY goTodoService /app

CMD [ "/app/goTodoService" ]

EXPOSE 8080