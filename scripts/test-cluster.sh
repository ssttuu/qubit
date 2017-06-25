#!/usr/bin/env bash

branch=${1:-}
hash=${2:-}
cluster_name=branch-${branch}
zone=us-east1-c

function cleanup() {
    gcloud -q container clusters delete ${cluster_name} --no-async --zone=${zone}
}
trap cleanup EXIT

gcloud -q container clusters create ${cluster_name} --no-async --zone=${zone} --num-nodes=1
gcloud -q container clusters get-credentials ${cluster_name} --zone=${zone}

./scripts/deploy.sh ${branch} ${hash}
