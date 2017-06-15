#!/usr/bin/env bash

set -euvo pipefail

TEST=${TEST:-false}
CIRCLE_SHA1=${CIRCLE_SHA1:-}

deployService() {
    set -euvo pipefail
    TEST=${TEST:-false}

    deploy_cmd="gcloud service-management deploy applications/${1}/services/api_config.yaml proto-gen/services/${1}/${1}.pb --format json"
    if [ "${TEST}" = "true" ]; then
      deploy_cmd="${deploy_cmd} --validate-only"
    fi
    deploy_response=$(${deploy_cmd})
    if [ $? != 0 ]; then
       exit 1
    fi
    echo ${deploy_response} | jq -r '.serviceConfig.id'
}

api_id=$(deployService "api")
compute_id=$(deployService "compute")

helm_cmd="helm upgrade --install --debug --recreate-pods --reset-values --wait --set Api.ApiId=${api_id},Compute.ApiId=${compute_id},Githash=${CIRCLE_SHA1} qubit ./helm/qubit/"
if [ "${TEST}" = "true" ]; then
  helm_cmd="${helm_cmd} --dry-run"
fi

helm init --client-only
${helm_cmd}
