#!/bin/bash
set -eu

tag=${1:-"latest"}
image_name="gcr.io/anan-project/research-db:${tag}"

build () {
    docker build --compress --pull -t "${image_name}" .
}

push () {
    docker push "$image_name"
}

deployment() {
    sed s/\${IMAGE_TAG}/${tag}/ deployment.yml >> kubernetes.yml
    echo "---" >> kubernetes.yml
    cat service.yml >> kubernetes.yml
}

build && deployment
