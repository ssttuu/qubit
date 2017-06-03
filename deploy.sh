#!/usr/bin/env bash

set -eo pipefail

# Deploy Compute Service
#echo "Deploying Compute to Google Endpoints"
#compute_service_response=$(gcloud service-management deploy compute/api_config.yaml compute/protos/compute/compute.pb --format json)
#
#compute_service_id=$(echo ${compute_service_response} | jq -r '.serviceConfig.id')
#compute_service_name=$(echo ${compute_service_response} | jq -r '.serviceConfig.name')

deployService() {
    echo $(gcloud service-management deploy services/${1}/api_config.yaml proto-gen/services/${1}.pb --format json) | jq -r '.serviceConfig.id'
}

api_id=$(deployService "api")

helm init --client-only

echo "Dry Run"
helm upgrade --install --dry-run --debug --recreate-pods --reset-values --wait \
    --set Api.ApiId=${api_id},Githash=${CIRCLE_SHA1} \
    qubit ./helm/qubit/

echo "Deploying"
helm upgrade --install --debug --recreate-pods --reset-values --wait \
    --set Api.ApiId=${api_id},Githash=${CIRCLE_SHA1} \
    qubit ./helm/qubit/
