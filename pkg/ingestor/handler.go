package ingestor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davidmontoyago/go-event-ingestor-api/pkg/payload"
)

// IngestPayload defers the reading/processing of the request payload to workers
func IngestPayload(w http.ResponseWriter, r *http.Request, workQueue *WorkQueue) {
	jsonPayload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, toJSONError(err), http.StatusBadRequest)
		return
	}

	workQueue.Enqueue(func() (payload.Payload, error) {
		return payload.FromJSON(jsonPayload)
	})

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func toJSONError(err error) string {
	errorResponse := new(bytes.Buffer)
	json.NewEncoder(errorResponse).Encode(map[string]string{"ok": "false", "error": err.Error()})
	return errorResponse.String()
}
