#!/usr/bin/env bash
set -eu

cwd=$(realpath $(dirname $0))

wait_db() {
  count=0
  ok=true
  while [ $count -lt 10 ];
  do
    [[ $(psql -U dolphin -d research -h 127.0.0.1 -c "select 1;") ]] && return 0
    sleep 5
    count=$(( count+1 ))
  done
  echo "failed to setup db" && return 1
}
run() {
  pushd ${cwd}/../apps/$1
  ./run.sh
  popd
}

run research-db && wait_db || exit 1
run research-api
run grpc-gateway
