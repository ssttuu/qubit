#!/usr/bin/env bash

set -euo pipefail

proto_name=${1}
proto_type=${2:-}
proto_path=./protos/${proto_name}.proto
base_protoc_command="/usr/bin/protoc -I/protobuf -I/googleapis -I./protos/"
protoc_command=

if [ -z "${proto_type}" ] || [ "${proto_type}" = "service" ]; then
	service_descriptor_path=./proto-gen/services/${proto_name}.pb
	mkdir -p $(dirname ${service_descriptor_path})
	protoc_command="${base_protoc_command} \
	    --include_imports \
	    --include_source_info \
	    ${proto_path} \
	    --descriptor_set_out \
	    ${service_descriptor_path}
	"
fi

if [ -z "${proto_type}" ] || [ "${proto_type}" = "go" ]; then
    modules=(
        google/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations
        geometry/geometry.proto=github.com/stupschwartz/qubit/proto-gen/go/geometry
        operators/operators.proto=github.com/stupschwartz/qubit/proto-gen/go/operators
        parameters/parameters.proto=github.com/stupschwartz/qubit/proto-gen/go/parameters
        images/images.proto=github.com/stupschwartz/qubit/proto-gen/go/images
    )
    module_string=
    for m in "${modules[@]}"; do
        module_string="${module_string}M${m},"
    done
	protoc_command="${protoc_command}${base_protoc_command} \
	    --go_out=${module_string}plugins=grpc:./proto-gen/go/ \
	    ${proto_path}
	"
fi

if [ -z "${proto_type}" ] || [ "${proto_type}" = "js" ]; then
	protoc_command="${protoc_command}${base_protoc_command} \
		--js_out=import_style=commonjs,binary:./proto-gen/js/ \
		--plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
		--grpc_out=./proto-gen/js/ \
		${proto_path}
	${base_protoc_command} \
		--js_out=import_style=commonjs,binary:./tests/integration/protos/ \
		--plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
		--grpc_out=./tests/integration/protos/ \
		${proto_path}
    "
fi

docker run --rm -v `pwd`:/workspace stupschwartz/protoman /bin/bash -c "${protoc_command}"
