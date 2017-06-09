#!/usr/bin/env bash

set -euo pipefail

CIRCLE_SHA1=${CIRCLE_SHA1:-}

deployService() {
    set -euo pipefail
    echo $(gcloud service-management deploy services/${1}/api_config.yaml proto-gen/services/${1}/${1}.pb --format json) | jq -r '.serviceConfig.id'
}

api_id=$(deployService "api")
compute_id=$(deployService "compute")

helm init --client-only

echo "Dry Run"
helm upgrade --install --dry-run --debug --recreate-pods --reset-values --wait \
    --set Api.ApiId=${api_id},Compute.ApiId=${compute_id},Githash=${CIRCLE_SHA1} \
    qubit ./helm/qubit/

echo "Deploying"
helm upgrade --install --debug --recreate-pods --reset-values --wait \
    --set Api.ApiId=${api_id},Compute.ApiId=${compute_id},Githash=${CIRCLE_SHA1} \
    qubit ./helm/qubit/
