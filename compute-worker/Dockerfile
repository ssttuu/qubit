FROM golang:1.8

ENV APP /go/src/github.com/stupschwartz/qubit/compute-worker

ADD . ${APP}
WORKDIR ${APP}
CMD go run main.go
