API_FILES = $(shell find services/api -type f -name "*.go")
COMPUTE_FILES = $(shell find services/compute -type f -name "*.go")
IMAGES_FILES = $(shell find services/images -type f -name "*.go")

# First target is default
build-go: build-api-go build-compute-go build-images-go

clean:
	touch services/api/run && rm services/api/run
	touch services/compute/run && rm services/compute/run
	touch services/images/run && rm services/images/run

# Go binaries
services/api/run: $(API_FILES)
	gofmt -l -s -w $(API_FILES) && cd services/api && go get ./... && go build -o run
build-api-go: services/api/run

services/compute/run: $(COMPUTE_FILES)
	gofmt -l -s -w $(COMPUTE_FILES) && cd services/compute && go get ./... && go build -o run
build-compute-go: services/compute/run

services/images/run: $(IMAGES_FILES)
	gofmt -l -s -w $(IMAGES_FILES) && cd services/images && go get ./... && go build -o run
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
