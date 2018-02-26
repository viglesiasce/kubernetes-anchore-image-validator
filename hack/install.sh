#!/bin/bash
RELEASE_NAME=${RELEASE_NAME:-analysis}
NAMESPACE=${NAMESPACE:-anchore}
pushd anchore-policy-validator/
  helm dep build
  helm install --timeout 600 --namespace $NAMESPACE -n $RELEASE_NAME .
popd
