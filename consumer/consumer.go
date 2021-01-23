package main

import (
	"github.com/dengpju/higo-rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	client := rabbitmq.New(rabbitmq.Host("192.168.42.131"))
	defer client.Close()

	go rabbitmq.Channel().
		Qos(1, 0, false).
		Consumer("usertest", "userreg", "c1", SendMail)
	//go rabbitmq.Consumer("usertest", "userreg", "c2", SendMail)
	//go rabbitmq.Consumer("usertestuion", "userreg", "c3", SendMail)
	//go rabbitmq.Consumer("usertestuion", "userreg", "c4", SendMail)
	select {}

}

func SendMail(msgs <-chan amqp.Delivery, cname string) {
	for msg := range msgs {
		go Send(msg, cname)
		/**
		{
			fmt.Printf("消费者:%s 消息id:%s 消息:%s \n", cname, msg.MessageId, string(msg.Body))
		}
		_ = msg.Ack(false) // 确认机制，如果不确认，服务停掉后，消息会从Unacked回到Ready中被其他消费者获取
		//time.Sleep(time.Second * 2)

		 */
	}
}

func Send(msg amqp.Delivery, cname string)  {
	time.Sleep(time.Second * 3)
	log.Printf("消费者:%s 消息id:%s 消息:%s \n", cname, msg.MessageId, string(msg.Body))
	_ = msg.Ack(false)
}