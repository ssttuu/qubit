FROM golang:1.8

ENV APP github.com/stupschwartz/qubit/compute

ADD . /go/src/${APP}
RUN go get ${APP}/...
RUN go install ${APP}
ENTRYPOINT /go/bin/server

EXPOSE 8080
