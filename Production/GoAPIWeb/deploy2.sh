#!/bin/bash

BUILD_NUMBER=$1
DOCKER_ID=$2
DOCKER_IMAGE=$3
DOCKER_PORTS=$4
# stop all running containers with our web application
docker stop `docker ps -a | grep ${DOCKER_ID}/${DOCKER_IMAGE} | awk '{print substr ($0, 0, 12)}'`
# remove all of those containers
docker rm `docker ps -a | grep ${DOCKER_ID}/${DOCKER_IMAGE} | awk '{print substr ($0, 0, 12)}'`
docker pull ${DOCKER_ID}/${DOCKER_IMAGE}:${BUILD_NUMBER}
docker run -d -ti -p ${DOCKER_PORTS} ${DOCKER_ID}/${DOCKER_IMAGE}:${BUILD_NUMBER}
