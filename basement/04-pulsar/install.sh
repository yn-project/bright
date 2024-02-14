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

helm repo add apache https://pulsar.apache.org/charts
helm repo update
helm install -n kube-system pulsar apache/pulsar --version 3.2.0

kubectl get pods -n kube-system | grep pulsar
