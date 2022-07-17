FROM golang:1.17.5-alpine3.14

RUN apk add --update alpine-sdk make git && \
    mkdir /mnt/workspace && \
    mkdir /workspace

COPY work /workspace

WORKDIR /tmp

# sql-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /workspace
