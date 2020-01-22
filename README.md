# go-event-ingestor-api

Event ingesting APIs can be useful in Event Sourcing architectures to provide an interface for event producers and prevent coupling to backend systems.

Uses a channel work queue and worker goroutines to concurrently read (fan-out) payloads and send them to the backend. Uses a proxy channel (ingest) and a context cancelation signal to ensure all workers finish before program shutdown.

Also, uses [codec](https://github.com/ugorji/go) for high performance JSON (un)marshalling.

## Run it
``` bash
make run
```

## Benchmark it

``` bash
./benchmark.sh
```

## Generate JSON codec

```
make codec
```