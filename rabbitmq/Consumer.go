package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type channel struct {
	channel *amqp.Channel
}

func Channel() *channel {
	c, err := Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &channel{c}
}

func (this *channel) Qos(prefetchCount, prefetchSize int, global bool) *channel {
	err := this.channel.Qos(prefetchCount, prefetchSize, global)
	if err != nil {
		log.Fatal(err)
	}
	return this
}

func (this *channel) Consumer(queue string, key string, cname string, callback func(msgs <-chan amqp.Delivery, cname string)) {
	defer this.channel.Close()
	msgs, err := this.channel.Consume(queue, key, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	callback(msgs, cname)
}
