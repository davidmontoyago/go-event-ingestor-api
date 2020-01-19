package main

import (
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
	router.HandleFunc("/ingest", IngestPayload).Methods("PUT")

	hostname, err := os.Hostname()
	if err != nil {
		log.Error.Fatal(err)
		return
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:8080", hostname),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Error.Fatal(srv.ListenAndServe())
}
