proto-deps:
	docker run --rm -v ${PWD}:/workspace --entrypoint bash stupschwartz/protoman -c 'protoc -I/protobuf -I/googleapis -I/googleapis/google/api/ --js_out=import_style=commonjs,binary:./tests/integration/protos/ /googleapis/google/api/*.proto'

compute-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./compute/protos/compute/ -I ./server/protos/ --include_imports --include_source_info ./compute/protos/compute/compute.proto --descriptor_set_out ./compute/protos/compute/compute.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./compute/protos/compute/ -I ./server/protos/ --go_out=plugins=grpc:./compute/protos/compute/ ./compute/protos/compute/compute.proto

server-geometry-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --include_imports --include_source_info ./server/protos/geometry/geometry.proto --descriptor_set_out ./server/protos/geometry/geometry.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:./server/protos/geometry/ ./server/protos/geometry/geometry.proto
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --js_out=import_style=commonjs,binary:./tests/integration/protos/ --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./tests/integration/protos/ ./server/protos/geometry/geometry.proto

server-images-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --include_imports --include_source_info ./server/protos/images/images.proto --descriptor_set_out ./server/protos/images/images.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,Mgeometry/geometry.proto=github.com/stupschwartz/qubit/server/protos/geometry,plugins=grpc:./server/protos/images/ ./server/protos/images/images.proto

server-health-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --include_imports --include_source_info ./server/protos/health/health.proto --descriptor_set_out ./server/protos/health/health.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:./server/protos/health/ ./server/protos/health/health.proto

server-organizations-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --include_imports --include_source_info ./server/protos/organizations/organizations.proto --descriptor_set_out ./server/protos/organizations/organizations.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:./server/protos/organizations/ ./server/protos/organizations/organizations.proto
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --js_out=import_style=commonjs,binary:./tests/integration/protos/ --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./tests/integration/protos/ ./server/protos/organizations/organizations.proto

server-scenes-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --include_imports --include_source_info ./server/protos/scenes/scenes.proto --descriptor_set_out ./server/protos/scenes/scenes.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:./server/protos/scenes/ ./server/protos/scenes/scenes.proto
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --js_out=import_style=commonjs,binary:./tests/integration/protos/ --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./tests/integration/protos/ ./server/protos/scenes/scenes.proto

server-operators-protos:
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --include_imports --include_source_info ./server/protos/operators/operators.proto --descriptor_set_out ./server/protos/operators/operators.pb
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,Mgeometry/geometry.proto=github.com/stupschwartz/qubit/server/protos/geometry,plugins=grpc:./server/protos/operators/ ./server/protos/operators/operators.proto
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/ --js_out=import_style=commonjs,binary:./tests/integration/protos/ --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./tests/integration/protos/ ./server/protos/operators/operators.proto

server-protos: server-geometry-protos server-images-protos server-health-protos server-organizations-protos server-scenes-protos server-operators-protos
	docker run --rm -v ${PWD}:/workspace stupschwartz/protoman -I ./server/protos/server/ -I ./server/protos/ --include_imports --include_source_info ./server/protos/server/server.proto --descriptor_set_out ./server/protos/server/server.pb

protos: proto-deps compute-protos server-protos


build-compute-go:
	cd compute && go get ./... && go build && cd ..

build-server-go:
	cd server && go get ./... && go build && cd ..

build-go: build-compute-go build-server-go


docker-build-go:
	docker run -it -v `pwd`:/go/src/github.com/stupschwartz/qubit -w /go/src/github.com/stupschwartz/qubit/compute golang:1.8 bash -c "go get ./...; go build; chmod 777 ./compute"
	docker run -it -v `pwd`:/go/src/github.com/stupschwartz/qubit -w /go/src/github.com/stupschwartz/qubit/server golang:1.8 bash -c "go get ./...; go build; chmod 777 ./server"

vendor:
	cd server && govendor update +external && govendor update +vendor && cd ..
	cd compute && govendor update +external && govendor update +vendor && cd ..

build-server:
	docker-compose build server

build-compute:
	docker-compose build compute

build: proto vendor build-server build-compute

up: build
	docker-compose up server compute
