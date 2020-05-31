#!/bin/bash
set -ue
selector="app=buildkitd"
context="minikube"

buildkitd_pods() {
  kubectl get pods -n buildkit \
    --selector=$selector \
    --field-selector=status.phase=Running \
    -o=go-template \
    --template="{{range .items}}{{.metadata.name}}{{\"\n\"}}{{end}}" \
    --sort-by="{.metadata.name}"
}

key=$1
if [ $# -ne 1 ]; then
  echo "require an argument: ./build.sh [pod_name]"
  exit 1
fi

which consistenthash
if [ $? -ne 1 ]; then
  echo "require consistenthash command install"
  exit 1
fi
echo "consistenthash command check ok"

pod=$(buildkitd_pods | consistenthash $key)

export BUILDKIT_HOST=