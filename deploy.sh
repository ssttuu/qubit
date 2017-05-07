#!/usr/bin/env bash

# Deploy Compute Service

yaml2json='python -c "import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)"'

service_response=$(gcloud service-management deploy compute/compute.pb compute/api_config.yaml --format json)

service_id=$(echo ${service_response} | jq '.serviceConfig.id')
service_name=$(echo ${service_response} | jq '.serviceConfig.name')

echo "${service_id} of ${service_name}"

bookstore_template_yaml=$(jinja2 grpc-bookstore.deployment.yaml -D id=${service_id} -D name=${service_name})

bookstore_template_json=$(echo "${bookstore_template_yaml}" | eval ${yaml2json})

k8_result=$(echo ${bookstore_template_json} | kubectl create -f -)


echo ${k8_result}

