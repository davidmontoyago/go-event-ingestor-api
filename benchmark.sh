#!/bin/sh
set -eu -o pipefail

HOST="localhost"
PORT="8080"
URL="http://$HOST:$PORT/ingest"

REQUESTS=1000000
CONCURRENCY=100

BECHNMARK_VERSION=$(git rev-parse HEAD)

echo "load testing $URL..."

hey -h2 -n $REQUESTS -c $CONCURRENCY -m PUT -H "Content-Type: application/json" -D payloads/event.json $URL > ./benchmarks/$BECHNMARK_VERSION.test

cat ./benchmarks/$BECHNMARK_VERSION.test
