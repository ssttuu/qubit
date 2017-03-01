FROM golang:1.8

ENV APP /go/src/github.com/stupschwartz/qubit/server

ADD . ${APP}
WORKDIR ${APP}
CMD go run main.go

EXPOSE 8443