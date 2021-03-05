package queue

type Queue interface {
	connect(map[string]string)
	publish(string)
	consume() string
}
