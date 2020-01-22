//go:generate codecgen -o payload_generated.go payload.go

package payload

import "github.com/ugorji/go/codec"

// Payload event type
type Payload struct {
	CorrelationID   string
	OriginTimestamp string
	IsSynthetic     bool `codec:",omitempty"`
	ApplicationID   string
	APIVersion      string `codec:"apiVersion"`
	EventHeader     string
	EventBody       string
}

// Reader deserializes a json payload... json string must be passed in as part of a closure
type Reader func() (Payload, error)

// FromJSON deserializes a JSON payload using codecs
func FromJSON(jsonPayload []byte) (Payload, error) {
	handler := new(codec.JsonHandle)
	var payload Payload
	err := codec.NewDecoderBytes(jsonPayload, handler).Decode(&payload)
	return payload, err
}
