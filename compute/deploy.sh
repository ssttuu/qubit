#!/usr/bin/env bash

yaml2json='python -c "import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)"'

# Deploy Compute Service
service_response=$(gcloud service-management deploy compute.pb api_config.yaml --format json)

service_id=$(echo ${service_response} | jq -r '.serviceConfig.id')
service_name=$(echo ${service_response} | jq -r '.serviceConfig.name')

echo "Deployment: ${service_id} of ${service_name}"

# Evaluate jinja templated YAML file
compute_template_yaml=$(jinja2 k8s/compute.deployment.yaml -D id=${service_id} -D name=${service_name} -D githash=$(git rev-parse HEAD))

# Convert YAML to JSON
compute_template_json=$(echo "${compute_template_yaml}" | eval ${yaml2json})

# Create deployment using JSON configuration
k8_result=$(echo ${compute_template_json} | kubectl create -f -)


echo ${k8_result}

