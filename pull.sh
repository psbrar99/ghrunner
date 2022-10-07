#!/usr/bin/env bash

set -o errexit
set -o pipefail
# set -x

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

IMAGES=(app
operator
pilot
proxyv2)

IMAGE_SUFFIXES=("debug" "distroless")

for image in "${IMAGES[@]}"; do
  for suffix in "${IMAGE_SUFFIXES[@]}"; do
    echo $HUB/${image}:${TAG}-${suffix} >> list.txt
    docker pull $HUB/${image}:${TAG}-${suffix}
  done
done