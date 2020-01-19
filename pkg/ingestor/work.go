package ingestor

import (
	"github.com/davidmontoyago/go-event-ingestor-api/pkg/backend"
	"github.com/davidmontoyago/go-event-ingestor-api/pkg/payload"
)

// WorkQueue contains a channel
type WorkQueue struct {
	queue      chan payload.PayloadReader
	maxWorkers int
}

// NewWorkQueue initializes worker queue
func NewWorkQueue(maxQueue int, maxWorkers int) *WorkQueue {
	return &WorkQueue{
		queue:      make(chan payload.PayloadReader, maxQueue),
		maxWorkers: maxWorkers,
	}
}

// StartWorkProcessor reads off the work queue and defers to the backend
func (w WorkQueue) StartWorkProcessor() {
	go func() {
		for payload := range w.queue {
			backend.Process(payload)
		}
	}()
}

// StartWorkProcessorPool fires a set of goroutines that will be reading work off the channel
func (w WorkQueue) StartWorkProcessorPool() {
	for i := 0; i < w.maxWorkers; i++ {
		w.StartWorkProcessor()
	}
}

// Enqueue adds a payload to the processing queue
func (w WorkQueue) Enqueue(payload payload.PayloadReader) {
	w.queue <- payload
}
