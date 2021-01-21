package main

import (
	"fmt"
	"github.com/dengpju/higo-rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func main() {
	client := rabbitmq.New(rabbitmq.Host("192.168.8.99"))
	defer client.Close()

	messageId := 1

	for {
		rand.Seed(time.Now().Unix())
		rabbitmq.Mq(rabbitmq.NewQueues().Append(rabbitmq.Queue("usertest"),
			rabbitmq.Queue("usertestuion")),
			rabbitmq.Exchange("UserExchange", "direct"),
			"userreg",
		).Message(amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   fmt.Sprintf("%d",messageId),
			Body:        []byte(fmt.Sprintf("gggggg%d", rand.Intn(1000)+1)),
		}).
			Message(amqp.Publishing{
				ContentType: "text/plain",
				MessageId:   fmt.Sprintf("%d",messageId + 1),
				Body:        []byte(fmt.Sprintf("hhhhhh%d", rand.Intn(1000)+1)),
			}).
			Message(amqp.Publishing{
				ContentType: "text/plain",
				MessageId:   fmt.Sprintf("%d",messageId + 2),
				Body:        []byte(fmt.Sprintf("jjjjjj%d", rand.Intn(1000)+1)),
			}).
			Send()
		messageId += 3
		log.Println("发送消息成功")
		time.Sleep(time.Second * 3)
	}

}
