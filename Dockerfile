FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /user_management
RUN mkdir -p /user_management/usermgmt

WORKDIR /user_management

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./usermgmt/* ./usermgmt/
COPY ./usermgmt_server/server.go .

RUN go build -o /user_management

EXPOSE 50051

CMD [ "./user_management" ]