package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/davidmontoyago/go-event-ingestor-api/pkg/ingestor"
	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/gorilla/mux"
)

func main() {
	// start work queue and workers
	maxQueue := getEnvAsIntOrFail("MAX_QUEUE")
	maxWorkers := getEnvAsIntOrFail("MAX_WORKERS")
	workQueue := ingestor.NewWorkQueue(maxQueue, maxWorkers)
	var waitgroup sync.WaitGroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	workQueue.StartWorkProcessorPool(ctx, &waitgroup)

	// do graceful termination
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		cancelFunc()
		waitgroup.Wait()
		log.Info.Println("all workers finished... exiting now.")
		os.Exit(0)
	}()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		ingestor.IngestPayload(w, r, workQueue)
	}).Methods("PUT")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Error.Fatal(srv.ListenAndServe())
}

func getEnvAsIntOrFail(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Error.Fatal(fmt.Sprintf("must specify %s: %v", key, err))
	}
	return value
}
