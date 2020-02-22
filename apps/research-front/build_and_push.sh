#!/bin/sh

registry=research-front
tag=$(git rev-parse --short HEAD | cut -c 1-8)

docker build -t ${registry}:${tag} .
docker build -t ${registry}:latest .

docker push ${registry}:${tag}
docker push ${registry}:latest
