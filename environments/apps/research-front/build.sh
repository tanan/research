#!/usr/bin/env bash
set -eu

DOCKER_BUILDKIT=1
export DOCKER_BUILDKIT

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
}

push() {
    docker push ${name}:${tag}
}

#copy && build && push && clean
copy && build && clean