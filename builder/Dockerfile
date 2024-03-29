FROM ubuntu:16.04

RUN apt-get update -y
RUN apt-get install -y build-essential lsb-release curl apt-transport-https ssh software-properties-common zip sudo
RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

RUN echo "deb https://packages.cloud.google.com/apt cloud-sdk-$(lsb_release -c -s) main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN apt-add-repository 'deb https://apt.dockerproject.org/repo ubuntu-xenial main'

RUN apt-get update -y

RUN apt-get install -y google-cloud-sdk

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl

RUN apt-cache policy docker-engine
RUN apt-get install -y docker-engine

RUN curl -L https://github.com/docker/compose/releases/download/1.14.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose

RUN apt-get install -y kubectl
RUN apt-get install -y jq

RUN curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
RUN unzip protoc-3.3.0-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/bin/
RUN mv protoc3/include/* /usr/local/include/

RUN curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get | bash

# migrate
RUN curl -L https://packagecloud.io/mattes/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/mattes/migrate/ubuntu/ xenial main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update -y
RUN apt-get install -y migrate
