#!/usr/bin/env bash

set -euo pipefail

for f in $(find protos -type f -name "*.proto"); do
    proto_path=$(echo "${f}" | awk '{split($0,a,"."); print a[1]}' | cut -c 8-)
    echo "Generating all protos for ${proto_path}"
    ${PWD}/scripts/gen-proto.sh "${proto_path}"
done
