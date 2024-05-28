#! /bin/bash
### Script to run a golang docker image

set -ex

readonly app_name=reporter
readonly docker_tag=v1.0.0

### Deploy docker container
deploy_app (){
    local app=${1}
    local tag=${2}

    echo "Deploying container image..."
    stop_container ${app}
    docker run -d --name ${app} -p ${API_PORT}:${API_PORT} \
        -e "API_PORT=${API_PORT}" \
        -e "EMAIL_ADDRESS=${EMAIL_ADDRESS}" \
        -e "EMAIL_PASSWORD=${EMAIL_PASSWORD}" \
        -e "EMAIL_PASSWORD=${EMAIL_PASSWORD}" \
        -e "GIN_MODE=release" \
        local/${app}:${tag}
}

### Stop docker container
stop_container() {
    local name=${1}

    docker stop ${name} || true 
    docker rm ${name} || true
}

main() {
    echo "Run deploy commands"

    deploy_app ${app_name} ${docker_tag}

    echo "Success!"
}

main