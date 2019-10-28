package ingestor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
)

var workQueue *WorkQueue

func init() {
	maxQueue := getEnvAsIntOrFail("MAX_QUEUE")
	maxWorkers := getEnvAsIntOrFail("MAX_WORKERS")
	workQueue = NewWorkQueue(maxQueue, maxWorkers)
	workQueue.StartWorkProcessorPool()
}

func IngestPayload(w http.ResponseWriter, r *http.Request) {
	jsonPayload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, ToJsonError(err), http.StatusBadRequest)
		return
	}

	workQueue.Enqueue(func() (Payload, error) {
		return FromJson(jsonPayload)
	})

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func ToJsonError(err error) string {
	errorResponse := new(bytes.Buffer)
	json.NewEncoder(errorResponse).Encode(map[string]string{"ok": "false", "error": err.Error()})
	return errorResponse.String()
}

func getEnvAsIntOrFail(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Error.Fatal(fmt.Sprintf("must specify %s: %v", key, err))
	}
	return value
}
