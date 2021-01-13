package main

import (
	"github.com/dengpju/higo-rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

func main()  {
	client :=rabbitmq.New()
	defer client.Close()
	channel, err := client.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	queue, err := channel.QueueDeclare("test", false, false,false,false,nil)
	if err != nil {
		log.Fatal(err)
	}
	err = channel.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte("test002"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("发送消息成功")
}
