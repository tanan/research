#!/bin/sh

tag=latest

docker stop research-front
docker rm -v research-front
docker run --network research -d -p 80:80 -v $(pwd)/environment/conf.d:/etc/nginx/conf.d:ro --name research-front research-front:${tag}
