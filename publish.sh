#!/bin/bash

echo tag?
read TAG
echo $TAG

# build dockerfile and push to dockerhub
docker build -t japiirainen/go-oluet-api:$TAG .
docker push japiirainen/go-oluet-api:$TAG

# deploy changes on personal vps
ssh personal "docker pull japiirainen/go-oluet-api:$TAG && docker tag japiirainen/go-oluet-api:$TAG dokku/oluet-api:$TAG && dokku tags:deploy oluet-api $TAG"
