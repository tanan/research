#!/bin/bash
set -eu

tag=${1:-"latest"}
image_name="gcr.io/anan-project/research-db"

build () {
    docker build --compress --pull -t "${image_name}:${tag}" .
    docker tag ${image_name}:${tag} ${image_name}:latest
}

deployment() {
    sed s/\${IMAGE_TAG}/${tag}/ deployment.yml >> kubernetes.yml
    echo "---" >> kubernetes.yml
    cat service.yml >> kubernetes.yml
}

build && deployment
