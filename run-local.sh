#!/bin/bash
docker_image=$(docker build -q .)

if [ -z ${docker_image} ]; then
    echo "${docker_image}"
    echo "DOCKER IMAGE COULD NOT BE BUILT"
else
    echo "RUNNING DOCKER IMAGE: ${docker_image}"
    docker run -it --rm \
        -p 9090:9090 \
        -e DATAINJESTOR_CONFIG_FILE="/etc/datainjestor/config.yml" \
        -v $(pwd)/config.yml:/etc/datainjestor/config.yml:ro \
        ${docker_image} $@
fi