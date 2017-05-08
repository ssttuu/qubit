#!/usr/bin/env bash

set -euo pipefail

echo "Update apt-get"
apt-get update -y
apt-get install -y lsb-release curl apt-transport-https ssh software-properties-common
apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

echo "Set CLOUD_SDK_REPO"
export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)"

echo "Update sources list"
echo "deb https://packages.cloud.google.com/apt ${CLOUD_SDK_REPO} main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
apt-add-repository 'deb https://apt.dockerproject.org/repo ubuntu-xenial main'


echo "Update apt-get again"
apt-get update -y

echo "Install Google Cloud SDK"
apt-get install -y google-cloud-sdk

echo "Install Kubectl"
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl

echo "Install Docker"
apt-cache policy docker-engine
apt-get install -y docker-engine

