package ingestor

import (
	"fmt"
	"time"

	log "github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
)

func Process(reader PayloadReader) {
	payload, error := reader()
	if error != nil {
		log.Error.Printf("failed to ready payload %v", error)
		// log to an error/broken letter  queue
	}
	time.Sleep(100 * time.Millisecond)
	log.Info.Println(fmt.Sprintf("ingested event for application %s [CorrelationId: %s]",
		payload.ApplicationId, payload.CorrelationId))
}
