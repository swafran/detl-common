package queue

type Queuer interface {
	connect(map[string]string)
	publish(string)
	consume() string
}
