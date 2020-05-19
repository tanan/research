#!/bin/bash

name=grpc-gateway-svc
network=research
registry="gcr.io/anan-project"

count=$(docker network ls | grep ${network} | wc -l)

[[ ${count} -ne 1 ]] && docker network create ${network}

docker stop ${name} && docker rm ${name}
docker pull ${registry}/grpc-gateway:latest
docker run --network ${network} --name ${name} -d -p 5000:5000 ${registry}/grpc-gateway:latest
