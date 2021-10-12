package queues

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	detl "github.com/swafran/detl-common"
	"github.com/swafran/detl-transform/handlers"
)

const (
	// TODO decide on handling credentials, but temporarily, for now:
	user = "guest"
	pass = "guest"
)

//RabbitQueue is a communication service to rabbitmq
type RabbitQueue struct {
	URL           string
	ReadQueue     string
	WriteExchange string
	WriteKey      string
	Handler       handlers.Handler
	Conn          *amqp.Connection
}

//Init establishes connection to queue server
func (q *RabbitQueue) Init(conf map[string]string) {
	var err error
	q.Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", user, pass, q.URL))
	detl.FailOnError(err, "Failed connection")
}

//Close closes amqp connection
func (q *RabbitQueue) Close() {
	q.Conn.Close()
}

//Consume one message from queue
func (q *RabbitQueue) Consume() {
	ch, _ := q.Conn.Channel()
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"extractedq", // name of the queue

		//TODO set to false
		true, // should the message be persistent? also queue will survive if the cluster gets reset
		//

		false, // autodelete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false, // exclusive means I should get an error if any other consumer subsribes to this queue
		false, // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		nil,   // arguments for more advanced configuration
	)

	detl.FailOnError(err, "Rabbit: failed to declare queue")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	detl.FailOnError(err, "Rabbit: failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			q.Handler.Handle(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	//q.Handler.Handle(q, msgs)

	//return "{doy:{joey{name:'joey', age:45, pets:['fifi', 'roxie', 'loulou']}}}"
}

//Publish message to queue
func (q *RabbitQueue) Publish(m string) {

}
