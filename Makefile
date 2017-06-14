API_FILES = $(shell find applications/api/services -type f -name "*.go")
COMPUTE_FILES = $(shell find applications/compute/services -type f -name "*.go")

# First target is default
build-go: fmt build-api-go build-compute-go

clean:
	touch applications/api/services/run && rm applications/api/services/run
	touch applications/compute/services/run && rm applications/compute/services/run

fmt:
	go fmt ./...

# Go binaries
applications/api/services/run: $(API_FILES)
	cd applications/api/services && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
build-api-go: applications/api/services/run

applications/compute/services/run: $(COMPUTE_FILES)
	cd applications/compute/services && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
build-compute-go: applications/compute/services/run

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

protonames = $(shell find protos -name "*.proto" | xargs -n1 basename | awk '{split($$0,a,"."); print a[1]}')

protos: $(foreach protoname,$(protonames),$(subst NAME,$(protoname),proto-gen/services/NAME/NAME.pb proto-gen/go/NAME/NAME.pb.go proto-gen/js/NAME/NAME_pb.js))

# Run containers
up: build
	docker-compose up server compute

test:
	./run-tests.integration.sh

migrate-cockroachdb:
	migrate -database
