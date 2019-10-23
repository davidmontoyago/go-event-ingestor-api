#!/bin/sh
set -eu -o pipefail

HOST="$(hostname)"
PORT="8080"

REQUESTS=1000000
CONCURRENCY=1000

BECHNMARK_VERSION=$(git rev-parse HEAD)

docker run --rm -v `pwd`:`pwd` -w `pwd` -p 8080:8080 --network="bridge" jordi/ab \
  -k -n $REQUESTS -c $CONCURRENCY \
  -T "Content-Type: application/json" \
  -u payloads/event.json "http://$HOST:$PORT/ingest" > ./benchmarks/$BECHNMARK_VERSION.test

cat ./benchmarks/$BECHNMARK_VERSION.test
