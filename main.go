package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/davidmontoyago/go-event-ingestor-api/pkg/ingestor"
	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingest", ingestPayload).Methods("PUT")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Error.Fatal(srv.ListenAndServe())
}

func ingestPayload(w http.ResponseWriter, r *http.Request) {
	payload, err := FromJson(r)
	if err != nil {
		http.Error(w, ToJsonError(err), http.StatusBadRequest)
		return
	}

	go Dispatch(payload)

	log.Info.Println(fmt.Sprintf("ingested event for application %s [CorrelationId: %s]", payload.ApplicationId,
		payload.CorrelationId))
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
