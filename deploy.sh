#!/usr/bin/env bash

set -eo pipefail

# Deploy Compute Service
echo "Deploying to Google Endpoints"
service_response=$(gcloud service-management deploy compute/protos/compute/compute.pb compute/api_config.yaml --format json)

service_id=$(echo ${service_response} | jq -r '.serviceConfig.id')
service_name=$(echo ${service_response} | jq -r '.serviceConfig.name')

helm init
helm install --debug --set Compute.ApiId=${service_id},Githash=${CIRCLE_SHA1} ./qubit-helm
