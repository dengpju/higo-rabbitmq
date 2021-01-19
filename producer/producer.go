package main

import (
	"github.com/dengpju/higo-rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

func main()  {
	client :=rabbitmq.New(rabbitmq.Host("192.168.8.99"))
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
			Body:[]byte("test005"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("发送消息成功")
	rabbitmq.Mq(rabbitmq.NewQueues().Append(rabbitmq.Queue("usertest"),
		rabbitmq.Queue("usertestuion")),
		rabbitmq.Exchange("UserExchange", "direct"),
		"userreg",
	).Message("gggggg").Message("kekekek").Message("hhhh").Send()
}
