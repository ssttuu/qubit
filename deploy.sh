#!/usr/bin/env bash

set -eo pipefail

# Deploy Compute Service
echo "Deploying Compute to Google Endpoints"
compute_service_response=$(gcloud service-management deploy compute/protos/compute/compute.pb compute/api_config.yaml --format json)

compute_service_id=$(echo ${compute_service_response} | jq -r '.serviceConfig.id')
compute_service_name=$(echo ${compute_service_response} | jq -r '.serviceConfig.name')

echo "Deploying Server to Google Endpoints"
server_service_response=$(gcloud service-management deploy server/protos/health/health.pb server/api_config.yaml --format json)

server_service_id=$(echo ${server_service_response} | jq -r '.serviceConfig.id')
server_service_name=$(echo ${server_service_response} | jq -r '.serviceConfig.name')

helm init --client-only

echo "Dry Run"
helm upgrade --install --dry-run --debug --recreate-pods --reset-values --wait --set Compute.ApiId=${compute_service_id},Server.ApiId=${server_service_id},Githash=${CIRCLE_SHA1} qubit ./helm/qubit/

echo "Deploying"
helm upgrade --install --debug --recreate-pods --reset-values --wait --set Compute.ApiId=${compute_service_id},Server.ApiId=${server_service_id},Githash=${CIRCLE_SHA1} qubit ./helm/qubit/
