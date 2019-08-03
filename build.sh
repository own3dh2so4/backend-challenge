#!/bin/bash

USER_ID=`id -u`
GROUP_ID=`id -g`
docker-compose run -u ${USER_ID}:${GROUP_ID} grpc_builder -f basket_service.proto -l go -o proto/basket_service
docker-compose run go_builder go build -a -installsuffix cgo .
docker-compose run go_builder chown ${USER_ID}:${GROUP_ID} prueba_cabify
