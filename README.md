# go-event-ingestor-api

Event ingesting APIs can be useful in Event Sourcing architectures to provide an interface for event producers and prevent coupling to backend systems.

Uses a channel work queue and worker goroutines to concurrently read (fan-out) payloads and send them to the backend.

## Run it
``` bash
make run
```

## Benchmark it

``` bash
./benchmark.sh
```
