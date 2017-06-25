APP_FILES = $(shell find applications -type f -name "*.go")
LIB_FILES = $(shell find applications/lib -type f -name "*.go")
API_SERVICE_FILES = $(shell find applications/api/services -type f -name "*.go")
API_TASK_FILES = $(shell find applications/api/tasks -type f -name "*.go")
COMPUTE_SERVICE_FILES = $(shell find applications/compute/services -type f -name "*.go")
COMPUTE_TASK_FILES = $(shell find applications/compute/tasks -type f -name "*.go")
CORE_FILES = $(shell find core -type f -name "*.go")
PROTO_FILES = $(shell find protos -type f -name "*.proto")

# First target is default
build-go: fmt bindata build-api-go build-compute-go

clean:
	touch applications/api/services/web/run && rm applications/api/services/web/run
	touch applications/compute/services/web/run && rm applications/compute/services/web/run
	touch applications/compute/services/coordinator/run && rm applications/compute/services/coordinator/run
	touch applications/compute/services/processor/run && rm applications/compute/services/processor/run

configure:
	go get -u github.com/jteeuwen/go-bindata/...

.fmt: $(APP_FILES) $(CORE_FILES)
	go fmt ./applications/...
	go fmt ./core/...
	touch .fmt
fmt: .fmt

#############
# Go binaries
#############

applications/api/tasks/migrate/migrations/bindata.go: $(shell find applications/api/tasks/migrate/migrations/sql -name "*.sql")
	# Using -prefix "sql/" because go-bindata postgres filename parsing can't handle path prefix
	cd applications/api/tasks/migrate/migrations && go-bindata -pkg migrations -prefix "sql/" sql

applications/compute/tasks/migrate/migrations/bindata.go: $(shell find applications/compute/tasks/migrate/migrations/sql -name "*.sql")
	# Using -prefix "sql/" because go-bindata postgres filename parsing can't handle path prefix
	cd applications/compute/tasks/migrate/migrations && go-bindata -pkg migrations -prefix "sql/" sql

bindata: \
	applications/api/tasks/migrate/migrations/bindata.go \
	applications/compute/tasks/migrate/migrations/bindata.go

applications/api/services/web/run: $(API_SERVICE_FILES) $(LIB_FILES) $(CORE_FILES)
	cd applications/api/services/web && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
applications/api/tasks/migrate/run: $(API_TASK_FILES) $(LIB_FILES)
	cd applications/api/tasks/migrate && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
build-api-go: \
	fmt \
	applications/api/services/web/run \
	applications/api/tasks/migrate/run

applications/compute/services/web/run: $(COMPUTE_SERVICE_FILES) $(LIB_FILES) $(CORE_FILES)
	cd applications/compute/services/web && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
applications/compute/services/coordinator/run: $(COMPUTE_SERVICE_FILES) $(LIB_FILES) $(CORE_FILES)
	cd applications/compute/services/coordinator && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
applications/compute/services/processor/run: $(COMPUTE_SERVICE_FILES) $(LIB_FILES) $(CORE_FILES)
	cd applications/compute/services/processor && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
applications/compute/tasks/migrate/run: $(COMPUTE_TASK_FILES) $(LIB_FILES)
	cd applications/compute/tasks/migrate && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
build-compute-go: \
	fmt \
	applications/compute/services/web/run \
	applications/compute/services/coordinator/run \
	applications/compute/services/processor/run \
	applications/compute/tasks/migrate/run

###############
# Docker images
###############

build-api: all-protos build-api-go
	docker-compose build api-web

build-compute: all-protos build-compute-go
	docker-compose build compute-web

build: build-api build-compute

.PHONY: bootstrap-postgres
bootstrap-postgres: build-api
	./scripts/bootstrap-postgres.sh

#####################
# Generate proto code
#####################

.PHONY: protoman
protoman:
	docker build -t us.gcr.io/qubit-161916/protoman ./protoman

all-protos: $(PROTO_FILES)
	rm -rf proto-gen || :
	mkdir -p proto-gen/services
	mkdir -p proto-gen/go
	mkdir -p proto-gen/js
	./scripts/generate-protos.sh

proto-gen/services/%.pb: protos/%.proto
	./scripts/generate-protos.sh $* service

proto-gen/go/%.pb.go: protos/%.proto
	./scripts/generate-protos.sh $* go

proto-gen/js/%_pb.js proto-gen/js/%_grpc_pb.js: protos/%.proto
	./scripts/generate-protos.sh $* js

tests/integration/protos/%_pb.js tests/integration/protos/%_grpc_pb.js: protos/%.proto
	./scripts/generate-protos.sh $* js-test

protonames = $(shell find protos -type f -name "*.proto" | xargs -n1 basename | awk '{split($$0,a,"."); print a[1]}')

protos: $(foreach protoname,$(protonames),$(subst NAME,$(protoname),proto-gen/services/NAME/NAME.pb proto-gen/go/NAME/NAME.pb.go proto-gen/js/NAME/NAME_pb.js tests/integration/protos/NAME/NAME_pb.js))

################
# Run containers
################

up: build
	docker-compose up server compute

.PHONY: test
test:
	./run-tests.integration.sh

migrate-cockroachdb:
	migrate -database
