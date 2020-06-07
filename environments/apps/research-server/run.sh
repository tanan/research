#!/bin/bash

name=research-api-svc
network=research
registry="gcr.io/anan-project"

count=$(docker network ls | grep ${network} | wc -l)

[[ ${count} -ne 1 ]] && docker network create ${network}

docker stop ${name} && docker rm ${name}
docker pull ${registry}/research-api:latest
docker run --network ${network} --name ${name} -d -p 19003:19003 ${registry}/research-api:latest
