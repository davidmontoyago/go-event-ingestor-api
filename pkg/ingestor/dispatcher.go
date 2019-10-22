package ingestor

import "time"

func Dispatch(payload *Payload) {
	time.Sleep(100 * time.Millisecond)
}
