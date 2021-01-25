package main

import (
	"fmt"
	"github.com/dengpju/higo-rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

var messageId int

func main() {
	client := rabbitmq.New(rabbitmq.Host("192.168.8.99"))
	defer client.Close()

	messageId = 1

	//http.HandleFunc("/producer", func(w http.ResponseWriter, r *http.Request) {
	//	mq := rabbitmq.Mq(rabbitmq.NewQueues().Append(rabbitmq.Queue("usertest"),
	//		/**rabbitmq.Queue("usertestuion")*/),
	//		rabbitmq.Exchange("UserExchange", "direct"),
	//		"userreg",
	//	).Confirm(true).NotifyReturn()
	//
	//	rand.Seed(time.Now().Unix())
	//	mq.Message(amqp.Publishing{
	//		ContentType: "text/plain",
	//		MessageId:   fmt.Sprintf("%d", messageId),
	//		Body:        []byte(fmt.Sprintf("gggggg%d", rand.Intn(1000)+1)),
	//	}).Send()
	//	messageId += 1
	//	log.Println("发送消息成功")
	//
	//	w.Header().Set("Content-type", "application/json")
	//	w.Write([]byte(fmt.Sprintf(`{"code":"200","data":%d}`, messageId)))
	//})
	//log.Println("服务启动成功")
	//log.Fatal(http.ListenAndServe(":8080", nil))
	mq := rabbitmq.Mq().
		DeclareExchange(rabbitmq.Exchange("UserExchange", "direct")).
		DeclareBindQueue(rabbitmq.NewQueues().Append(rabbitmq.Queue("usertest"),
			/**rabbitmq.Queue("usertestuion")*/), "userreg")

	for {
		mq.SetConfirm(true).NotifyReturn().SetExchange(rabbitmq.Exchange("UserExchange", "direct")).SetKey("userreg")
		rand.Seed(time.Now().Unix())
		mq.Context(amqp.Publishing{
			//Headers: map[string]interface{}{"x-delay":3000},// 延迟
			ContentType: "text/plain",
			MessageId:   fmt.Sprintf("%d", messageId),
			Body:        []byte(fmt.Sprintf("gggggg%d", rand.Intn(1000)+1)),
		}).Send()
		messageId += 1
		log.Println("发送消息成功")
		time.Sleep(time.Second * 10)
		mq = rabbitmq.Mq()
	}

}
