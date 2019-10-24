package ingestor

type WorkQueue struct {
	queue chan Payload
}

func NewWorkQueue() *WorkQueue {
	return &WorkQueue{
		queue: make(chan Payload),
	}
}

// Reads of the work queue and defers to the backend
func (w WorkQueue) StartWorkProcessor() {
	for {
		select {
		case payload := <-w.queue:
			go Process(payload)
		}
	}
}

// Fires a set of goroutines that will be reading work of the queue channel
func (w WorkQueue) StartWorkProcessorPool(maxProcessors int) {
	for i := 0; i < maxProcessors; i++ {
		go w.StartWorkProcessor()
	}
}

func (w WorkQueue) Enqueue(payload Payload) {
	w.queue <- payload
}
