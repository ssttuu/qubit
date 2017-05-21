#!/usr/bin/env bash

set -eo pipefail

# Create Compute Service Description
docker run --rm -v $(pwd):/workspace -w /workspace znly/protoc --include_imports --include_source_info compute/compute.proto --descriptor_set_out compute/compute.pb

# Deploy Compute Service
echo "Deploying to Google Endpoints"
service_response=$(gcloud service-management deploy compute/compute.pb compute/api_config.yaml --format json)

service_id=$(echo ${service_response} | jq -r '.serviceConfig.id')
service_name=$(echo ${service_response} | jq -r '.serviceConfig.name')

helm init
helm install --dry-run --debug --set Compute.ApiId=${service_id},Githash=${CIRCLE_SHA1} ./qubit-helm
