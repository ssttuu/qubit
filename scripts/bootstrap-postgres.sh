#!/usr/bin/env bash

set -euo pipefail

postgres_data_dir=/var/lib/postgresql/data
postgres_host=postgres.api.qubit.site
postgres_url=postgres://postgres@${postgres_host}:5432/postgres?sslmode=disable

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

wait_for_postgres_script='
host='"${postgres_host}"'
port=5432
limit=15
start=$(date +%s)
until [ "$(PGCONNECT_TIMEOUT=1 psql -U postgres -h ${host} -t -p ${port} -c "select now()" postgres 2>/dev/null || :)" ]; do
    sleep 1
    now=$(date +%s)
    if [ "$((now - start - limit))" -gt "0" ]; then
        (>&2 echo "${host} is not accessible after ${limit} seconds")
        exit 1
    fi
done
'

echo "Waiting for postgres"
docker run --rm \
    --network qubit-api-postgres-bootstrap \
    postgres \
    /bin/bash -c "${wait_for_postgres_script}"

# TODO: Don't use default "postgres" user and DB
commands='
    /app/run migrate up
'

echo "Running migrations"
docker run --rm \
    --volume ${PWD}/applications/api/tasks:/app \
    --network qubit-api-postgres-bootstrap \
    --env POSTGRES_URL=${postgres_url} \
    golang:1.8 \
    /bin/bash -c "${commands}"
