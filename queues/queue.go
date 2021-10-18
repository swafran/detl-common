package queues

//Queue interacts with queues for messaging between stages of ETL
type Queue interface {
	Init(map[string]string)
	Publish(string)
	Consume()
	Close()
}

//Handler processes queue messages
type Handler interface {
	Handle(string)
}
