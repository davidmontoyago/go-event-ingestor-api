package producer

type Payload struct {
	CorrelationId  string
	OriginDatetime string
	IsSynthetic    bool
	ApplicationId  string
	EventHeader    string
	EventBody      string
}
