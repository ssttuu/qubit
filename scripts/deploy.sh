#!/usr/bin/env bash

set -euvo pipefail

branch=${1:-}
hash=${2:-}

env=branch
if [ "${branch}" = "master" ]; then
    env=production
fi
values_path=values.${env}.yaml

deployService() {
    set -euvo pipefail
    app=${1}
    svc=${2}

    deploy_response=$(gcloud service-management deploy applications/${app}/services/${svc}/api_config.yaml proto-gen/services/${app}/${app}.pb --format json)
    if [ $? != 0 ]; then
       exit 1
    fi
    echo ${deploy_response} | jq -r '.serviceConfig.id'
}

api_web_id=$(deployService "api" "web")
compute_web_id=$(deployService "compute" "web")

helm init --client-only
helm upgrade --install --debug --recreate-pods --reset-values --wait -f ${values_path} \
    --set Api.ApiId=${api_web_id},Compute.ApiId=${compute_web_id},Githash=${hash} qubit ./helm/qubit/
