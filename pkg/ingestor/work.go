package ingestor

type WorkQueue struct {
	queue chan Payload
}

func NewWorkQueue() *WorkQueue {
	return &WorkQueue{
		queue: make(chan Payload, 10000),
	}
}

// Reads off the work queue and defers to the backend
func (w WorkQueue) StartWorkProcessor() {
	go func() {
		for payload := range w.queue {
			Process(payload)
		}
	}()
}

// Fires a set of goroutines that will be reading work off the channel
func (w WorkQueue) StartWorkProcessorPool(maxProcessors int) {
	for i := 0; i < maxProcessors; i++ {
		w.StartWorkProcessor()
	}
}

func (w WorkQueue) Enqueue(payload Payload) {
	w.queue <- payload
}
