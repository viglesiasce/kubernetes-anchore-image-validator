#!/bin/bash
RELEASE_NAME=${RELEASE_NAME:-analysis}
NAMESPACE=${NAMESPACE:-anchore}
pushd anchore-policy-validator/
  helm dep build
  helm install --namespace $NAMESPACE -n $RELEASE_NAME .
popd