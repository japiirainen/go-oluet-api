#!/bin/bash

echo tag?
read TAG
echo $TAG

docker build -t japiirainen/go-oluet-api:$TAG .
docker push japiirainen/go-oluet-api:$TAG
ssh personal "docker pull japiirainen/go-oluet-api:$TAG && docker tag japiirainen/go-oluet-api:$TAG dokku/oluet-api:$TAG && dokku tags:deploy oluet-api $TAG"