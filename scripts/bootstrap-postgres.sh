#!/usr/bin/env bash

set -euo pipefail

function bootstrap() {
    set -euo pipefail
    app=${1}

    postgres_data_dir=/var/lib/postgresql/data
    postgres_host=postgres.${app}.qubit.site
    postgres_url=postgres://postgres@${postgres_host}:5432/postgres?sslmode=disable

    docker volume rm qubit-${app}-postgres || :
    docker volume create qubit-${app}-postgres
    docker network create qubit-${app}-postgres-bootstrap 2>/dev/null || :
    container=$(docker run -d \
        --volume qubit-${app}-postgres:${postgres_data_dir} \
        --network qubit-${app}-postgres-bootstrap \
        --network-alias postgres.${app}.qubit.site \
        postgres)

    function cleanup() {
        docker stop ${container} || :
        docker rm ${container} || :
        docker network rm qubit-${app}-postgres-bootstrap || :
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

    echo "Waiting for ${app} postgres"
    docker run --rm \
        --network qubit-${app}-postgres-bootstrap \
        postgres \
        /bin/bash -c "${wait_for_postgres_script}"

    # TODO: Don't use default "postgres" user and DB
    commands='
        /app/run migrate up
    '

    echo "Running migrations"
    docker run --rm \
        --volume ${PWD}/applications/${app}/tasks/migrate:/app \
        --network qubit-${app}-postgres-bootstrap \
        --env POSTGRES_URL=${postgres_url} \
        golang:1.8 \
        /bin/bash -c "${commands}"

    docker stop ${container} || :
    docker rm ${container} || :
    docker network rm qubit-${app}-postgres-bootstrap || :
}

bootstrap api &
api_bootstrap_pid=$!

bootstrap compute &
compute_bootstrap_pid=$!

wait ${api_bootstrap_pid}
wait ${compute_bootstrap_pid}
