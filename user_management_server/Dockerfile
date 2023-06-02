FROM golang:alpine as server_env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /user_management
RUN mkdir -p /user_management/usermgmt

WORKDIR /user_management

ADD ../go.mod .
ADD ../go.sum .
RUN go mod download

COPY ../usermgmt/* ./usermgmt/
COPY ../user_management_server/server.go .

RUN go build -o /user_management_sever

EXPOSE 50051

CMD [ "./user_management_server" ]