#!/bin/bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
PROJECT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)



OrginazeName=bright

user=$(whoami)
service_name=jenkins
version=latest
if [ "$user" == "root" ]; then
    docker push ${registry}/${OrginazeName}/$service_name:$version
else
    sudo docker push ${registry}/${OrginazeName}/$service_name:$version
fi
