#!/usr/bin/env bash

set -euvo pipefail

postgres_data_dir=/var/lib/postgresql/data
postgres_url=postgres://postgres@postgres.api.qubit.site/postgres?sslmode=disable

docker volume rm qubit-api-postgres || :
docker volume create qubit-api-postgres
docker network create qubit-api-postgres-bootstrap
container=$(docker run -d \
    --volume qubit-api-postgres:${postgres_data_dir} \
    --network qubit-api-postgres-bootstrap \
    --network-alias postgres.api.qubit.site \
    postgres)

function cleanup() {
    docker stop ${container} || :
    docker rm ${container} || :
    docker network rm qubit-api-postgres-bootstrap || :
}
trap 'cleanup' EXIT

sleep 5

# TODO: Don't use default "postgres" user and DB
commands='
    /app/run migrate up
'

docker run --rm \
    --volume ${PWD}/applications/api/tasks:/app \
    --network qubit-api-postgres-bootstrap \
    --env POSTGRES_URL=${postgres_url} \
    golang:1.8 \
    /bin/bash -c "${commands}"
