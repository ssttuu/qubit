#!/usr/bin/env bash

set -euo pipefail

docker run --rm -v ${PWD}:/workspace --entrypoint bash stupschwartz/protoman -c 'protoc -I/protobuf -I/googleapis -I/googleapis/google/api/ --js_out=import_style=commonjs,binary:./proto-gen/js/ /googleapis/google/api/*.proto'
docker run --rm -v ${PWD}:/workspace --entrypoint bash stupschwartz/protoman -c 'protoc -I/protobuf -I/googleapis -I/googleapis/google/api/ --js_out=import_style=commonjs,binary:./tests/integration/protos/ /googleapis/google/api/*.proto'


genProto () {
    echo "Generating ${1}"
	docker run --rm -v `pwd`:/workspace stupschwartz/protoman \
	    -I ./protos/ \
	    --include_imports \
	    --include_source_info \
	    ./protos/${1}/${1}.proto \
	    --descriptor_set_out ./proto-gen/services/${1}.pb

	docker run --rm -v `pwd`:/workspace stupschwartz/protoman \
	    -I ./protos/ \
	    --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,Mgeometry/geometry.proto=github.com/stupschwartz/qubit/proto-gen/go/geometry,plugins=grpc:./proto-gen/go/ \
	    ./protos/${1}/${1}.proto

	docker run --rm -v `pwd`:/workspace stupschwartz/protoman \
	    -I ./protos/ \
	    --js_out=import_style=commonjs,binary:./proto-gen/js/ \
	    --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
	    --grpc_out=./proto-gen/js/ \
	    ./protos/${1}/${1}.proto

	docker run --rm -v `pwd`:/workspace stupschwartz/protoman \
	    -I ./protos/ \
	    --js_out=import_style=commonjs,binary:./tests/integration/protos/ \
	    --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
	    --grpc_out=./tests/integration/protos/ \
	    ./protos/${1}/${1}.proto
}

genProto "geometry"
genProto "health"
genProto "images"
genProto "operators"
genProto "organizations"
genProto "parameters"
genProto "scenes"
genProto "api"
