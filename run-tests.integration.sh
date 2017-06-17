#!/usr/bin/env bash

set -euvo pipefail

compose_args="-f docker-compose.yml -f docker-compose.test.integration.yml"
make bootstrap-postgres

function cleanup() {
    docker volume rm qubit-api-postgres || :
    docker-compose ${compose_args} logs || :
    docker-compose ${compose_args} down || :
}
trap 'cleanup' EXIT

docker-compose ${compose_args} build test
docker-compose ${compose_args} run test
