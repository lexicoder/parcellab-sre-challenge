#!/bin/bash

exit_on_failure() {
    echo "Exiting"
    exit 1
}

# check if docker is installed
docker version 1> /dev/null
if [ $? -ne 0 ]; then
    echo "docker is not installed!"
    exit_on_failure
fi

if [[ $DEPLOY_TARGET -eq "minikube" ]] || [[ $DEPLOY_TARGET -eq "kubernetes" ]]; then
    # check if kind is installed
    kind version 1> /dev/null
    if [ $? -ne 0 ]; then
        echo "minikube is not installed!"
        exit_on_failure
    fi

    # check if helm is installed
    helm version 1> /dev/null
    if [ $? -ne 0 ]; then
        echo "helm is not installed!"
        exit_on_failure
    fi
fi

exit 0
