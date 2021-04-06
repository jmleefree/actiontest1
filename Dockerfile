FROM golang:1.15.3 AS builder

ENV GO111MODULE on

ADD . /actiontest1

WORKDIR /actiontest1

RUN GOOS=linux go build -tags actiontest1 -o actiontest1 -v