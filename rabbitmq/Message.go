package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func Send(queue string, message string) error {
	channel, err := Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	q, err := channel.QueueDeclare(queue, false, false,false,false,nil)
	if err != nil {
		log.Fatal(err)
	}
	return channel.Publish("", q.Name, false, false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		},
	)
}
