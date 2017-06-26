#!/usr/bin/env bash

branch=${1:-}
prefix=${2:-}

found=0

for f in $(git diff --name-only origin/master ${branch}); do
    if [[ "$(expr substr ${f} 1 ${#prefix})" = "${prefix}" ]]; then
        found=$((${found}+1))
    fi
done

echo ${found}
