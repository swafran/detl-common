package queues

import (
	"fmt"

	"github.com/streadway/amqp"
	"gitlab.com/detl/detl-common"
)

const (
	// TODO decide on handling credentials, but temporarily, for now:
	user = "guest"
	pass = "guest"
)

//RabbitQueue is a communication service to rabbitmq
type RabbitQueue struct {
	URL           string
	ReadExchange  string
	ReadKey       string
	WriteExchange string
	WriteKey      string
	Conn          *amqp.Connection
}

//Init set fields and establishes connection
func (q *RabbitQueue) Init(conf map[string]string) {
	q.URL = conf["url"]
	q.ReadExchange = conf["readExchange"]
	q.ReadKey = conf["readKey"]
	q.WriteExchange = conf["writeExchange"]
	q.WriteKey = conf["writeKey"]

	var err error
	q.Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", user, pass, q.URL))
	detl.FailOnError(err, "Failed connection")
}

//Close closes amqp connection
func (q *RabbitQueue) Close() {
	q.Conn.Close()
}
