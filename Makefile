API_FILES = $(shell find services/api -type f -name "*.go")
COMPUTE_FILES = $(shell find services/compute -type f -name "*.go")
IMAGES_FILES = $(shell find services/images -type f -name "*.go")

# First target is default
build-go: build-api-go build-compute-go

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

# Docker images
build-api: build-api-go
	docker-compose build server

build-compute: build-compute-go
	docker-compose build compute

build: proto build-api build-compute

# Generate proto code
proto-gen/services/%.pb: protos/%.proto
	./scripts/gen-proto.sh service $*

proto-gen/go/%.pb.go: protos/%.proto
	./scripts/gen-proto.sh go $*

proto-gen/js/%_pb.js proto-gen/js/%_grpc_pb.js: protos/%.proto
	./scripts/gen-proto.sh js $*

protonames = geometry health images operators organizations parameters scenes api compute

protos: $(foreach protoname,$(protonames),$(subst NAME,$(protoname),proto-gen/services/NAME/NAME.pb proto-gen/go/NAME/NAME.pb.go proto-gen/js/NAME/NAME_pb.js))

# Run containers
up: build
	docker-compose up server compute

test:
	./run-tests.integration.sh

migrate-cockroachdb:
	migrate -database
