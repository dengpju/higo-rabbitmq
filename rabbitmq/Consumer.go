package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func Consumer(queue string, key string, callback func(<-chan amqp.Delivery)) {
	channel, err := Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	msgs, err := channel.Consume(queue, key, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	callback(msgs)
}
