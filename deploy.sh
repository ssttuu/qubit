#!/usr/bin/env bash

set -euvo pipefail

TEST=${TEST:-false}
CIRCLE_SHA1=${CIRCLE_SHA1:-}

deployService() {
    set -euvo pipefail
    TEST=${TEST:-false}
    app=${1}
    svc=${2}

    deploy_cmd="gcloud service-management deploy applications/${app}/services/${svc}/api_config.yaml proto-gen/services/${app}-${svc}/${app}-${svc}.pb --format json"
    if [ "${TEST}" = "true" ]; then
      deploy_cmd="${deploy_cmd} --validate-only"
    fi
    deploy_response=$(${deploy_cmd})
    if [ $? != 0 ]; then
       exit 1
    fi
    echo ${deploy_response} | jq -r '.serviceConfig.id'
}

api_web_id=$(deployService "api" "web")
compute_web_id=$(deployService "compute" "web")

helm_cmd="helm upgrade --install --debug --recreate-pods --reset-values --wait --set Api.ApiId=${api_web_id},Compute.ApiId=${compute_web_id},Githash=${CIRCLE_SHA1} qubit ./helm/qubit/"
if [ "${TEST}" = "true" ]; then
  helm_cmd="${helm_cmd} --dry-run"
fi

helm init --client-only
${helm_cmd}
