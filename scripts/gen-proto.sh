#!/usr/bin/env bash

set -euo pipefail

type=${1}
proto_name=${2}
proto_path=./protos/${proto_name}.proto
protoc_command="docker run --rm -v `pwd`:/workspace stupschwartz/protoman -I ./protos/"

if [ "${type}" = "service" ]; then
	service_descriptor_path=./proto-gen/services/${proto_name}.pb
	mkdir -p $(dirname ${service_descriptor_path})
	${protoc_command} \
		--include_imports \
		--include_source_info \
		${proto_path} \
		--descriptor_set_out ${service_descriptor_path}
elif [ "${type}" = "go" ]; then
	${protoc_command} \
		--go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,Mgeometry/geometry.proto=github.com/stupschwartz/qubit/proto-gen/go/geometry,Moperators/operators.proto=github.com/stupschwartz/qubit/proto-gen/go/operators,Mparameters/parameters.proto=github.com/stupschwartz/qubit/proto-gen/go/parameters,Mimages/images.proto=github.com/stupschwartz/qubit/proto-gen/go/images,plugins=grpc:./proto-gen/go/ \
		${proto_path}
elif [ "${type}" = "js" ]; then
	${protoc_command} \
		--js_out=import_style=commonjs,binary:./proto-gen/js/ \
		--plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
		--grpc_out=./proto-gen/js/ \
		${proto_path}
	${protoc_command} \
		--js_out=import_style=commonjs,binary:./tests/integration/protos/ \
		--plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
		--grpc_out=./tests/integration/protos/ \
		${proto_path}
fi

