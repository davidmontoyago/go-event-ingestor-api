package backend

import (
	"time"

	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/davidmontoyago/go-event-ingestor-api/pkg/payload"
)

// Process reads the payload and processes it
func Process(reader payload.PayloadReader) {
	payload, err := reader()
	if err != nil {
		log.Error.Println("failed to ready payload", err)
		// log to an error/broken letter  queue
		return
	}
	time.Sleep(100 * time.Millisecond)
	log.Info.Printf("ingested event for application %s [CorrelationId: %s]\n", payload.ApplicationId, payload.CorrelationId)
}
