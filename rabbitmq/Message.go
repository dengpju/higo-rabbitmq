package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type mq struct {
	channel  *amqp.Channel
	exchange *exchange
	key      string //路由键
	message  []string
}

type Queues []*queue

func NewQueues() Queues {
	return make(Queues, 0)
}

func (this Queues) Append(q ...*queue) Queues {
	this = append(this, q...)
	return this
}

type queue struct {
	Name string
}

func Queue(name string) *queue {
	return &queue{Name: name}
}

type exchange struct {
	Name string
	Kind string // direct、fanout、headers
}

func Exchange(name string, kind string) *exchange {
	return &exchange{Name: name, Kind: kind}
}

func Mq(queues Queues, exchange *exchange, key string) *mq {
	channel, err := Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	//声明交换机
	err = channel.ExchangeDeclare(exchange.Name, exchange.Kind, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 声明队列
	for _, queue := range queues {
		q, err := channel.QueueDeclare(queue.Name, false, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}
		// 绑定队列
		err = channel.QueueBind(q.Name, key, exchange.Name, false, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &mq{channel: channel, exchange: exchange, key: key}
}

func (this *mq) Message(context string) *mq {
	this.message = append(this.message, context)
	return this
}

func (this *mq) Send() (err []error) {
	defer this.channel.Close()
	if len(this.message) > 0 {
		for _,message :=range this.message {
			err = append(err, this.channel.Publish(this.exchange.Name, this.key, false, false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(message),
				},
			))
		}
	}
	return
}
