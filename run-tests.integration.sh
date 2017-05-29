#!/usr/bin/env bash

docker-compose -f docker-compose.test.integration.yml run test
docker-compose -f docker-compose.test.integration.yml down
