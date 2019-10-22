package ingestor

type Payload struct {
	CorrelationId   string
	OriginTimestamp string
	IsSynthetic     bool
	ApplicationId   string
	ApiVersion      string
	EventHeader     string
	EventBody       string
}
