#!/bin/bash

user="developdockerstate"
repo="zetrarepo"
name="datainjestor"
version="local"
latest=false

while test $# -gt 0; do
    case "$1" in
      -v|--version)
        shift
        if test $# -gt 0; then
            version=$1
        else
            echo "no specified version number"
            exit 1
        fi
        shift
        ;;
      --latest)
        shift
        latest=true
        ;;
      *)
        break
        ;;
    esac
done


if [ $version == "local" ]; then
    echo "Usage: .\build-publish.sh --version <version-number> (optional: --latest)"
    echo " "
    echo "--version (REQUIRED): Specifies the version number for images to push"
    echo "--latest (OPTIONAL): Inforces that image is pushed as latest as well"
    echo " "
    exit 1
fi

docker_image="$user/$repo"

echo "Building docker image..."
docker build -t $docker_image:$name-$version .

if [ $? -eq 0 ]; then
    echo "Pushing docker image..."
    docker push $docker_image:$name-$version
    if [ $? -ne 0 ]; then
        echo "Failed to push docker image"
        exit 1
    fi

    if [ $latest ]; then
        echo "Adding latest tag..."
        docker tag $docker_image:$name-$version $docker_image:$name-latest
        echo "Pushing docker image as latest..."
        docker push $docker_image:$name-latest
    fi
else
    echo "Failed to build image"
    exit 1
fi

# add support for helm charts in the future