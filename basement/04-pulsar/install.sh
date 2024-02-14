#!/bin/bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
PROJECT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)

set -o errexit
set -o nounset
set -o pipefail

helm install -n kube-system pulsar ./pulsar -f $SHELL_FOLDER/value.yaml
kubectl get pods -n kube-system | grep pulsar
