#!/usr/bin/env bash

set -eo pipefail

# Deploy Compute Service
echo "Deploying to Google Endpoints"
service_response=$(gcloud service-management deploy compute/protos/compute/compute.pb compute/api_config.yaml --format json)

service_id=$(echo ${service_response} | jq -r '.serviceConfig.id')
service_name=$(echo ${service_response} | jq -r '.serviceConfig.name')

helm init --client-only

echo "Dry Run"
helm upgrade --install --dry-run --debug --recreate-pods --reset-values --wait --set Compute.ApiId=${service_id},Githash=${CIRCLE_SHA1} qubit ./helm/qubit/

echo "Deploying"
helm upgrade --install --debug --recreate-pods --reset-values --wait --set Compute.ApiId=${service_id},Githash=${CIRCLE_SHA1} qubit ./helm/qubit/
