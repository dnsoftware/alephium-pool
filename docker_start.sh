#!/bin/bash

docker pull ddaemon/alph_pooler:latest || exit 1
docker stop alph_pooler || true
docker rm alph_pooler || true
docker run -d --name alph_pooler -p 8090:8090 ddaemon/alph_pooler:latest || exit 1
