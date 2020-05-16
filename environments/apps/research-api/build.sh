#!/usr/bin/env bash
set -eu

DOCKER_BUILDKIT=1
export DOCKER_BUILDKIT

tag=${1:-"latest"}
name="research-api"
registry="gcr.io/anan-project"

cwd=$(realpath $(dirname $0))

build() {
    cd ${cwd}/../../../apps/${name}
    docker build --compress --pull -f ${cwd}/Dockerfile -t ${registry}/${name}:${tag} .
}

push() {
    docker push ${name}:${tag}
}

# build && push
build
