package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	. "github.com/davidmontoyago/go-event-ingestor-api/pkg/ingestor"
	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingest", ingestPayload).Methods("PUT")

	hostname, err := os.Hostname()
	if err != nil {
		log.Error.Fatal(err)
		return
	}

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:8080", hostname),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	WorkQueue = make(chan Payload)
	StartWorkProcessorPool(10)

	log.Error.Fatal(srv.ListenAndServe())
}

var WorkQueue chan Payload

func StartWorkProcessor() {
	for {
		select {
		case payload := <-WorkQueue:
			go Process(payload)
		}
	}
}

func StartWorkProcessorPool(maxProcessors int) {
	for i := 0; i < maxProcessors; i++ {
		go StartWorkProcessor()
	}
}

func ingestPayload(w http.ResponseWriter, r *http.Request) {
	payload, err := FromJson(r)
	if err != nil {
		http.Error(w, ToJsonError(err), http.StatusBadRequest)
		return
	}

	WorkQueue <- *payload

	log.Info.Println(fmt.Sprintf("ingested event for application %s [CorrelationId: %s]", payload.ApplicationId,
		payload.CorrelationId))
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
