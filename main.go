package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/davidmontoyago/go-kafka-ingestor-api/pkg/kafka/log"
	producer "github.com/davidmontoyago/go-kafka-ingestor-api/pkg/kafka/producer"
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

	log.Fatal(srv.ListenAndServe())
}

func ingestPayload(w http.ResponseWriter, r *http.Request) {
	payload, err := fromJson(r)
	if err != nil {
		http.Error(w, toJsonError(err), http.StatusBadRequest)
		return
	}
	Info.Println(fmt.Sprintf("ingested event for application %s [CorrelationId: %s]", payload.ApplicationId,
		payload.CorrelationId))
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func fromJson(r *http.Request) (*producer.Payload, error) {
	decoder := json.NewDecoder(r.Body)
	var payload *producer.Payload
	err := decoder.Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func toJsonError(err error) string {
	errorResponse := new(bytes.Buffer)
	json.NewEncoder(errorResponse).Encode(map[string]string{"ok": "false", "error": err.Error()})
	return errorResponse.String()
}
