#!/usr/bin/env bash

echo  "GOOS=linux go build"
 GOOS=linux go build -o sidecar-injector

export HOST=$1
export TAG=0.0.19
docker build -t ${HOST}/istio/sidecar-injector:${TAG} .

docker login -p Harbor12345 -u admin ${HOST}

docker push ${HOST}/istio/sidecar-injector:${TAG}

rm -rf sidecar-injector
