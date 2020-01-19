package payload

import (
	"encoding/json"
)

// Payload event type
type Payload struct {
	CorrelationId   string
	OriginTimestamp string
	IsSynthetic     bool
	ApplicationId   string
	ApiVersion      string
	EventHeader     string
	EventBody       string
}

// PayloadReader deserializes a json payload... json string must be included in a closure
type PayloadReader func() (Payload, error)

// FromJSON deserializes a JSON payload
func FromJSON(jsonPayload []byte) (Payload, error) {
	payload := Payload{}
	err := json.Unmarshal(jsonPayload, &payload)
	return payload, err
}
