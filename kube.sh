#!/bin/bash

set -Eou pipefail

RESULT=$1?
deployApp(){
    kubectl apply -f books-config-k8s.yaml
    kubectl apply -f mysql-book-config.yml
    kubectl apply -f mysql-order-config.yml
    kubectl apply -f mysql-user-config.yml
    kubectl apply -f user-confing-k8s.yaml
    kubectl apply -f user-config-k8s.yaml
}

main() {
    deployApp
    if [ $RESULT == 0 ];
    then
        echo "success"
    else
        echo "failed"
    fi
}

main