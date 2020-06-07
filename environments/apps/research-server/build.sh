#!/usr/bin/env bash
set -eu

# DOCKER_BUILDKIT=1
# export DOCKER_BUILDKIT

tag=${1:-"latest"}
name="research-server"
registry="gcr.io/anan-project"

cwd=$(realpath $(dirname $0))

build() {
    cd ${cwd}/../../../apps/${name}
    ./gradlew build -x docker
    # docker build --compress --pull -f ${cwd}/Dockerfile -t ${registry}/${name}:${tag} .
    docker tag ${registry}/${name}:${tag} ${registry}/${name}:latest
}

push() {
    docker push ${name}:${tag}
}

deployment() {
    cd ${cwd}
    sed s/\${IMAGE_TAG}/${tag}/ deployment.yml >> kubernetes.yml
    echo "---" >> kubernetes.yml
    cat service.yml >> kubernetes.yml
}

# build && push
# build && deployment
build
