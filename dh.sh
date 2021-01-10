#!/bin/bash

echo tag?
read TAG
echo $TAG

docker build -t japiirainen/go-oluet-api:$TAG .
docker push japiirainen/go-oluet-api:$TAG