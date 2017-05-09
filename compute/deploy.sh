#!/usr/bin/env bash

set -euo pipefail

function onExit {
    popd
}
trap onExit EXIT
pushd `dirname $0`

# Create Compute Service Description
protoc -I /usr/local/include/ -I ./ --include_imports --include_source_info compute.proto --descriptor_set_out compute.pb

# Deploy Compute Service
service_response=$(gcloud service-management deploy compute.pb api_config.yaml --format json)

service_id=$(echo ${service_response} | jq -r '.serviceConfig.id')
service_name=$(echo ${service_response} | jq -r '.serviceConfig.name')

echo "Deployment: ${service_id} of ${service_name}"

deployment_formatted="k8s/compute.deployment.formatted.yaml"

# Evaluate jinja templated YAML file
jinja2 k8s/compute.deployment.yaml -D id=${service_id} -D name=${service_name} -D githash=$(git rev-parse HEAD) > ${deployment_formatted}

# Create deployment using JSON configuration
kubectl create -f ${deployment_formatted}
