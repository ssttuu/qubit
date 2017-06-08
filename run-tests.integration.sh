#!/usr/bin/env bash

set -euvo pipefail

compose_args="-f docker-compose.test.integration.yml"

function cleanup() {
    docker-compose ${compose_args} logs || :
    docker-compose ${compose_args} down || :
}
trap 'cleanup' EXIT

docker-compose ${compose_args} run test
