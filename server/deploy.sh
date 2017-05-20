#!/usr/bin/env bash

set -euo pipefail

function onExit {
    popd
}
trap onExit EXIT
pushd `dirname $0`

# Create Compute Service Description
protoc -I /usr/local/include/ -I ./ --include_imports --include_source_info server.proto --descriptor_set_out server.pb

# Deploy Compute Service
echo "Deploying to Google Endpoints"
service_response=$(gcloud service-management deploy compute.pb api_config.yaml --format json)

service_id=$(echo ${service_response} | jq -r '.serviceConfig.id')
service_name=$(echo ${service_response} | jq -r '.serviceConfig.name')

deployment_formatted="k8s/compute.deployment.formatted.yaml"

# Evaluate jinja templated YAML file
jinja2 k8s/compute.deployment.yaml -D id=${service_id} -D name=${service_name} -D githash=$(git rev-parse HEAD) > ${deployment_formatted}

# Create deployment using JSON configuration
echo "Deploying to Kubernetes"
kubectl create -f ${deployment_formatted}
