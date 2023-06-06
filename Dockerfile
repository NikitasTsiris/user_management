# Build Server Image
FROM golang:1.20.4 as server_builder

ENV GO111MODULE=on

RUN mkdir -p /usermgmt

COPY ./usermgmt/ /usermgmt/
COPY ./server/server.go /

COPY go.mod /
COPY go.sum /

RUN go mod download
RUN go mod tidy

WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch as server

COPY --from=server_builder /server /

EXPOSE 50051

CMD [ "/server" ]

# Build Client Image
FROM golang:1.20.4 as client_builder

ENV GO111MODULE=on

RUN mkdir -p /usermgmt

COPY ./usermgmt/ /usermgmt/
COPY ./client/client.go /

COPY go.mod /
COPY go.sum /

RUN go mod download
RUN go mod tidy

WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o client

FROM scratch as client

COPY --from=client_builder /client /

CMD [ "/client" ]