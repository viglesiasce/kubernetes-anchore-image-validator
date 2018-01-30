#!/bin/bash -xe
export DATE=$(date +%s)
export TAG=${VERSION:-dev-$DATE}
export REPO=${REPO:-viglesiasce/anchore-image-admission-server}
docker build --tag $REPO:$TAG .
docker tag $REPO:$TAG $REPO:latest
gcloud docker -- push $REPO:$TAG
gcloud docker -- push $REPO:latest
