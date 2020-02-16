#!/bin/bash
set -eu

tag=${1:-"latest"}
image_name="research-db:${tag}"

build () {
    docker build --compress --pull -t "${image_name}" .
}

push () {
    docker push "$image_name"
}

#build && push
build
