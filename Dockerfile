# Build Server Image
FROM golang:alpine as server_builder

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

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch as server

COPY --from=server_builder /server/server /

EXPOSE 50051

cmd [ "/server" ]

# Build Client Image
FROM golang:alpine as client_builder

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

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o client

FROM scratch as client

COPY --from=client_builder /client/client /


CMD [ "/client" ]