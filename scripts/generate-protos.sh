#!/usr/bin/env bash

set -euo pipefail

proto_name=${1:-}
proto_type=${2:-}
DOCKER=${DOCKER:-true}

base_protoc_command="/usr/bin/protoc -I/protobuf -I/googleapis -I./protos/"
protoc_command=

for f in $(find protos -type f -name "*.proto"); do
    proto=$(echo "${f}" | awk '{split($0,a,"."); print a[1]}' | cut -c 8-)
    if [ "${proto_name}" ] && [ "${proto_name}" != "${proto}" ]; then
        continue
    fi

    proto_path=./protos/${proto}.proto

    if [ -z "${proto_type}" ] || [ "${proto_type}" = "service" ]; then
        mkdir -p ./proto-gen/services
        service_descriptor_path=./proto-gen/services/${proto}.pb
        mkdir -p $(dirname ${service_descriptor_path})
        protoc_command="${protoc_command}${base_protoc_command} \
            --include_imports \
            --include_source_info \
            ${proto_path} \
            --descriptor_set_out \
            ${service_descriptor_path}
        "
    fi

    if [ -z "${proto_type}" ] || [ "${proto_type}" = "go" ]; then
        mkdir -p ./proto-gen/go
        modules=(
            computation_renders/computation_renders.proto=github.com/stupschwartz/qubit/proto-gen/go/computation_renders
            computations/computations.proto=github.com/stupschwartz/qubit/proto-gen/go/computations
            geometry/geometry.proto=github.com/stupschwartz/qubit/proto-gen/go/geometry
            google/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations
            image_sequences/image_sequences.proto=github.com/stupschwartz/qubit/proto-gen/go/image_sequences
            images/images.proto=github.com/stupschwartz/qubit/proto-gen/go/images
            organizations/organizations.proto=github.com/stupschwartz/qubit/proto-gen/go/organizations
            projects/projects.proto=github.com/stupschwartz/qubit/proto-gen/go/projects
            render_operators/render_operators.proto=github.com/stupschwartz/qubit/proto-gen/go/render_operators
            scene_events/scene_events.proto=github.com/stupschwartz/qubit/proto-gen/go/scene_events
            scene_renders/scene_renders.proto=github.com/stupschwartz/qubit/proto-gen/go/scene_renders
            scenes/scenes.proto=github.com/stupschwartz/qubit/proto-gen/go/scenes
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
        mkdir -p ./proto-gen/js
        protoc_command="${protoc_command}${base_protoc_command} \
            --js_out=import_style=commonjs,binary:./proto-gen/js/ \
            --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
            --grpc_out=./proto-gen/js/ \
            ${proto_path}
        "
    fi

    if [ -z "${proto_type}" ] || [ "${proto_type}" = "js-test" ]; then
        mkdir -p ./tests/integration/protos/
        protoc_command="${protoc_command}${base_protoc_command} \
            --js_out=import_style=commonjs,binary:./tests/integration/protos/ \
            --plugin=protoc-gen-grpc=/usr/lib/node_modules/grpc-tools/bin/grpc_node_plugin \
            --grpc_out=./tests/integration/protos/ \
            ${proto_path}
        "
    fi
done

if [ "${DOCKER}" = "true" ]; then
    docker run --rm -v ${PWD}:/workspace stuschwartz/protoman /bin/bash -c "${protoc_command}"
    docker run --rm -v ${PWD}:/workspace stuschwartz/protoman /bin/bash -c "chmod 777 -R proto-gen && chmod 777 -R tests/integration/protos"
else
    ${protoc_command}
    chmod 777 -R proto-gen && chmod 777 -R tests/integration/protos
fi

