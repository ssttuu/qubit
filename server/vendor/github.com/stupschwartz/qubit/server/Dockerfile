FROM golang:1.8

ENV APP /go/src/github.com/stupschwartz/qubit/server

RUN apt-get update -y && apt-get install -y ntp
RUN /etc/init.d/ntp restart

ADD . ${APP}
WORKDIR ${APP}
CMD go run main.go

EXPOSE 8443