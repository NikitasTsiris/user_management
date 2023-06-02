# Build Server Image
FROM golang:alpine as server

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /server
RUN mkdir -p /server/usermgmt

WORKDIR /server

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./usermgmt/ ./usermgmt/
COPY ./server/server.go .

RUN go build -o server

EXPOSE 50051

CMD [ "./server" ]

# Build Client Image
FROM golang:alpine as client

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /client
RUN mkdir -p /client/usermgmt

WORKDIR /client

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./usermgmt/ ./usermgmt/
COPY ./client/client.go .

RUN go build -o client

CMD [ "./client" ]