package ingestor

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func FromJson(r *http.Request) (*Payload, error) {
	decoder := json.NewDecoder(r.Body)
	var payload *Payload
	err := decoder.Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func ToJsonError(err error) string {
	errorResponse := new(bytes.Buffer)
	json.NewEncoder(errorResponse).Encode(map[string]string{"ok": "false", "error": err.Error()})
	return errorResponse.String()
}
