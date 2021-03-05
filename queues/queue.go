package queues

type Queue interface {
	Init(map[string]string)
	Publish(string)
	Consume() string
	Close()
}
