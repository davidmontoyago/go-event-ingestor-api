package ingestor

import (
	"context"
	"sync"

	"github.com/davidmontoyago/go-event-ingestor-api/pkg/backend"
	"github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/davidmontoyago/go-event-ingestor-api/pkg/payload"
)

// WorkQueue contains work queue channel for processing work and ingest channel to proxy for the queue channel
// and allow its graceful termination
type WorkQueue struct {
	queue      chan payload.Reader
	ingest     chan payload.Reader
	maxWorkers int
}

// NewWorkQueue initializes worker queue
func NewWorkQueue(maxQueue int, maxWorkers int) *WorkQueue {
	return &WorkQueue{
		queue:      make(chan payload.Reader, maxQueue),
		ingest:     make(chan payload.Reader, 1),
		maxWorkers: maxWorkers,
	}
}

// StartWorkProcessor reads off the work queue and defers to the backend
func (w WorkQueue) StartWorkProcessor(wg *sync.WaitGroup) {
	defer wg.Done()
	for payload := range w.queue {
		backend.Process(payload)
	}
}

// StartWorkProcessorPool fires a set of goroutines that will be reading work off the channel
// also, fires a function that serves as proxy for the work channel and allows graceful termination
func (w WorkQueue) StartWorkProcessorPool(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(w.maxWorkers)
	for i := 0; i < w.maxWorkers; i++ {
		go w.StartWorkProcessor(wg)
	}

	// Use proxy channel (w.ingest) to control access to the work queue and allow graceful termination
	// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	go func() {
		defer close(w.queue)
		for {
			select {
			case work := <-w.ingest:
				w.queue <- work
			case <-ctx.Done():
				log.Info.Println("received termination signal... closing work queue.")
				return
			}
		}
	}()
}

// Enqueue adds a payload to the processing queue
func (w WorkQueue) Enqueue(payload payload.Reader) {
	w.ingest <- payload
}
