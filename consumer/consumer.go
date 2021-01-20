package main

import (
	"fmt"
	"github.com/dengpju/higo-rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
)

func main()  {
	client :=rabbitmq.New(rabbitmq.Host("192.168.8.99"))
	defer client.Close()
	/**
	channel, err := client.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	msgs, err := channel.Consume("test","c1",false,false,false, false,nil)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range msgs{
		msg.Ack(true)// 确认机制，如果不确认，服务停掉后，消息会从Unacked回到Ready中被其他消费者获取
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}

	 */
	rabbitmq.Consumer("usertest", "userreg", SendMail)

}

func SendMail(msgs <-chan amqp.Delivery)  {
	for msg := range msgs {
		{
			fmt.Println(msg.DeliveryTag, string(msg.Body))
		}
		msg.Ack(false) // 确认机制，如果不确认，服务停掉后，消息会从Unacked回到Ready中被其他消费者获取
	}
}