package ingestor

import (
	"encoding/json"
)

type Payload struct {
	CorrelationId   string
	OriginTimestamp string
	IsSynthetic     bool
	ApplicationId   string
	ApiVersion      string
	EventHeader     string
	EventBody       string
}

type PayloadReader func() (Payload, error)

func FromJson(jsonPayload []byte) (Payload, error) {
	payload := Payload{}
	err := json.Unmarshal(jsonPayload, &payload)
	return payload, err
}
