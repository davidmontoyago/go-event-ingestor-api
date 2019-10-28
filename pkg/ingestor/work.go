package ingestor

type WorkQueue struct {
	queue      chan PayloadReader
	maxWorkers int
}

func NewWorkQueue(maxQueue int, maxWorkers int) *WorkQueue {
	return &WorkQueue{
		queue:      make(chan PayloadReader, maxQueue),
		maxWorkers: maxWorkers,
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
func (w WorkQueue) StartWorkProcessorPool() {
	for i := 0; i < w.maxWorkers; i++ {
		w.StartWorkProcessor()
	}
}

func (w WorkQueue) Enqueue(payload PayloadReader) {
	w.queue <- payload
}
