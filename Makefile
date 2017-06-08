# First target is default
build-go: build-api-go build-compute-go build-images-go

# Go binaries
services/api/run: $(shell find services/api -type f -name "*.go")
	cd services/api && go get ./... && go build -o run
build-api-go: services/api/run

services/compute/run: $(shell find services/compute -type f -name "*.go")
	cd services/compute && go get ./... && go build -o run
build-compute-go: services/compute/run

services/images/run: $(shell find services/images -type f -name "*.go")
	cd services/images && go get ./... && go build -o run
build-images-go: services/images/run

# Docker images
build-api: build-api-go
	docker-compose build server

build-compute: build-compute-go
	docker-compose build compute

build-images: build-images-go
	docker-compose build images

build: proto build-api build-compute build-images

# Generate proto code
$(shell find proto-gen) $(shell find tests/integration/protos): $(shell find protos -type f -name "*.proto")
	bash ./gen-proto.sh
proto: $(shell find proto-gen) $(shell find tests/integration/protos)

# Run containers
up: build
	docker-compose up server compute

test:
	./run-tests.integration.sh

migrate-cockroachdb:
	migrate -database
