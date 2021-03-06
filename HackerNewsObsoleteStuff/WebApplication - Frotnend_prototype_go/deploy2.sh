#!/bin/bash

BUILD_NUMBER=$1
DOCKER_ID=$2
# stop all running containers with our web application
docker stop `docker ps -a | grep ${DOCKER_ID}/gotestsite | awk '{print substr ($0, 0, 12)}'`
# remove all of those containers
docker rm `docker ps -a | grep ${DOCKER_ID}/gotestsite | awk '{print substr ($0, 0, 12)}'`
docker pull ${DOCKER_ID}/gotestsite:${BUILD_NUMBER}
docker run -d -ti -p 8080:8080 ${DOCKER_ID}/gotestsite:${BUILD_NUMBER}
