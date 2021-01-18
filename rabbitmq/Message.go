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
	q, err := channel.QueueDeclare(queue, false, false,false,false,nil)
	if err != nil {
		log.Fatal(err)
	}
	q2, err := channel.QueueDeclare(queue+"union", false, false,false,false,nil)
	if err != nil {
		log.Fatal(err)
	}
	err = channel.ExchangeDeclare("UserExchange", "direct",false,false, false, false, nil)
	if err != nil {
		return err
	}
	// bind
	err = channel.QueueBind(q.Name, "userreg", "UserExchange", false, nil)
	if err != nil {
		return err
	}
	err = channel.QueueBind(q2.Name, "userreg", "UserExchange", false, nil)
	if err != nil {
		return err
	}
	return channel.Publish("UserExchange", "userreg", false, false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		},
	)
}
