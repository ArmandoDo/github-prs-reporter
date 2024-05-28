#! /bin/bash
### Script to containerize a golang image

set -ex

readonly app_name=webui
readonly docker_tag=v1.0.0

### Build docker image
containerize_build (){
    local app=${1}
    local tag=${2}

    docker build --rm --no-cache --progress=plain \
        -t local/${app}:${tag} \
        -f docker/Dockerfile .
}

main() {
    echo "Containerize - Build"

    containerize_build ${app_name} ${docker_tag}

    echo "Success!"
}

main