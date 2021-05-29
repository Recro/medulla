#!/bin/bash

echo 'Generating cert-manager.yaml...'
curl https://github.com/jetstack/cert-manager/releases/download/v1.3.1/cert-manager.yaml -s -L -o cert-manager.yaml

echo 'Generating istio-operator.yaml...'
istioctl operator dump > istio-operator.yaml

echo 'Generating schemahero.yaml...'
kubectl schemahero install --yaml > schemahero.yaml

echo 'Generating tidb-crd.yaml...'
curl https://raw.githubusercontent.com/pingcap/tidb-operator/v1.1.12/manifests/crd.yaml -s -L -o tidb-crd.yaml

echo 'Generating tidb-operator.yaml...'
helm repo add pingcap https://charts.pingcap.org/ > /dev/null 2>&1
helm template --namespace tidb-operator tidb-operator pingcap/tidb-operator --version v1.1.12 > tidb-operator.yaml
