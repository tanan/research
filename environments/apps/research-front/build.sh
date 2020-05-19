#!/usr/bin/env bash
set -eu

# DOCKER_BUILDKIT=1
# export DOCKER_BUILDKIT

tag=${1:-"latest"}

cwd=$(realpath $(dirname $0))
src=sources
name="research-front"
registry="gcr.io/anan-project"

clean() {
    [[ -d $src && -w $src ]] && rm -rf $src
}

copy() {
    clean
    cp -R ${cwd}/../../../apps/${name} $src
}

build() {
    docker build --compress --pull -f ${cwd}/Dockerfile -t ${registry}/${name}:${tag} .
    docker tag ${registry}/${name}:${tag} ${registry}/${name}:latest
}

push() {
    docker push ${name}:${tag}
}

deployment() {
    cat configmap.yml >> kubernetes.yml
    echo "---" >> kubernetes.yml
    sed s/\${IMAGE_TAG}/${tag}/ deployment.yml >> kubernetes.yml
    echo "---" >> kubernetes.yml
    cat service.yml >> kubernetes.yml
}

#copy && build && push && clean
copy && build && clean && deployment