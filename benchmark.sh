#!/bin/sh
set -eu -o pipefail

HOST="127.0.0.1"
PORT="8080"

REQUESTS=100
CONCURRENCY=10

BECHNMARK_VERSION=$(git rev-parse HEAD)

ab -n $REQUESTS -c $CONCURRENCY \
   -T "Content-Type: application/json" -u payloads/event.json \
   "http://$HOST:$PORT/ingest" > ./benchmarks/$BECHNMARK_VERSION.test

cat ./benchmarks/$BECHNMARK_VERSION.test
