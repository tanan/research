#!/bin/bash

name=research-db-svc
network=research
registry="gcr.io/anan-project"

count=$(docker network ls | grep ${network} | wc -l)

[[ ${count} -ne 1 ]] && docker network create ${network}

docker rm -f ${name}
docker run --network ${network} -p 5432:5432 -d --name ${name} ${registry}/research-db:latest
