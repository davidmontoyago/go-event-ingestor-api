package main

import (
	"net/http"
	"time"

	"github.com/davidmontoyago/go-event-ingestor-api/pkg/ingestor"
	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingest", ingestor.IngestPayload).Methods("PUT")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Error.Fatal(srv.ListenAndServe())
}
