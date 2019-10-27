package ingestor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var workQueue *WorkQueue

func init() {
	workQueue = NewWorkQueue()
	workQueue.StartWorkProcessorPool(10000)
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
