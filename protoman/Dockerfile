FROM ubuntu:16.04

ENV GRPC_VERSION=1.0.1              \
    PROTOBUF_VERSION=3.1.0          \
    SWIFT_PROTOBUF_VERSION=0.9.24   \
    GOPATH=/go

RUN apt-get update -y
RUN apt-get install -y curl
RUN apt-get install -y sudo
RUN apt-get install -y automake
RUN apt-get install -y autoconf
RUN apt-get install -y libtool
RUN apt-get install -y git
RUN apt-get install -y golang-go
RUN apt-get install -y python-software-properties
RUN curl -sL https://deb.nodesource.com/setup_7.x | sudo -E bash -
RUN apt-get install -y nodejs
#RUN curl -L https://github.com/QuentinPerez/docker-alpine-swift-protobuf/releases/download/$SWIFT_PROTOBUF_VERSION/export-lib-056f6f6.tar | tar xv -C /
RUN curl -L https://github.com/google/protobuf/archive/v${PROTOBUF_VERSION}.tar.gz | tar xvz
#RUN curl -L https://raw.githubusercontent.com/QuentinPerez/docker-alpine-swift-protobuf/master/ressources/ld_library_path.patch > /ld_library_path.patch
RUN cd /protobuf-${PROTOBUF_VERSION} && autoreconf -f -i -Wall,no-obsolete && ./configure --prefix=/usr --enable-static=no && make && make install
#RUN rm /ld_library_path.patch
RUN git clone --recursive -b v${GRPC_VERSION} https://github.com/grpc/grpc.git
RUN cd /grpc/third_party/protobuf && git checkout v${PROTOBUF_VERSION}
RUN cd /grpc && \
        make plugins && make install-plugins prefix=/usr && \
        rm -rf `pwd`
RUN curl -L https://github.com/grpc/grpc-java/archive/v${GRPC_VERSION}.tar.gz | tar xvz
RUN go get -ldflags "-w -s" \
        github.com/golang/protobuf/protoc-gen-go \
        github.com/gogo/protobuf/protoc-gen-gofast \
        github.com/gogo/protobuf/protoc-gen-gogo \
        github.com/gogo/protobuf/protoc-gen-gogofast \
        github.com/gogo/protobuf/protoc-gen-gogofaster \
        github.com/gogo/protobuf/protoc-gen-gogoslick \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN install -c /go/bin/* /usr/bin/
RUN rm -rf /go/*
RUN mkdir -p /protobuf/google/protobuf
RUN for f in any duration descriptor empty struct timestamp wrappers; do \
            curl -L -o /protobuf/google/protobuf/${f}.proto https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/${f}.proto; \
        done
RUN git clone https://github.com/googleapis/googleapis.git
RUN mkdir -p /protobuf/github.com/gogo/protobuf/gogoproto
RUN curl -L -o /protobuf/github.com/gogo/protobuf/gogoproto/gogo.proto https://raw.githubusercontent.com/gogo/protobuf/master/gogoproto/gogo.proto

RUN npm install -g grpc-tools

RUN mkdir /workspace
WORKDIR /workspace

ENTRYPOINT ["/usr/bin/protoc", "-I/protobuf", "-I/googleapis"]