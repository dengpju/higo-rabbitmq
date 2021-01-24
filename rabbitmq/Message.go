package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type MQ struct {
	Channel       *amqp.Channel
	Exchange      *exchange
	Key           string //路由键
	Confirm       bool
	Message       amqp.Publishing
	notifyConfirm chan amqp.Confirmation
	notifyReturn  chan amqp.Return
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

func Mq() *MQ {
	channel, err := Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &MQ{Channel: channel}
}

func (this *MQ) DeclareExchange(exchange *exchange) *MQ {
	//声明交换机
	err := this.Channel.ExchangeDeclare(exchange.Name, exchange.Kind, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	this.Exchange = exchange
	return this
}

func (this *MQ) DeclareBindQueue(queues Queues, key string) *MQ {
	this.Key = key
	for _, queue := range queues {
		// 声明队列
		q, err := this.Channel.QueueDeclare(queue.Name, false, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}
		// 绑定队列
		err = this.Channel.QueueBind(q.Name, this.Key, this.Exchange.Name, false, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	return this
}

func (this *MQ) Context(context amqp.Publishing) *MQ {
	this.Message = context
	return this
}

func (this *MQ) Qos(prefetchCount, prefetchSize int, global bool) *MQ {
	err := this.Channel.Qos(prefetchCount, prefetchSize, global)
	if err != nil {
		log.Fatal(err)
	}
	return this
}

func (this *MQ) SetConfirm(b bool) *MQ {
	err := this.Channel.Confirm(false)
	if err != nil {
		log.Fatal(err)
	}
	this.notifyConfirm = this.Channel.NotifyPublish(make(chan amqp.Confirmation))
	this.Confirm = b
	return this
}

func (this *MQ) listenConfirm() {
	defer this.Channel.Close()
	ret := <-this.notifyConfirm
	if ret.Ack {
		log.Println("Confirm消息发送成功")
	} else {
		log.Println("Confirm消息发送失败")
	}
}

func (this *MQ) NotifyReturn() *MQ {
	this.notifyReturn = this.Channel.NotifyReturn(make(chan amqp.Return))
	go this.listenReturn()
	return this
}

func (this *MQ) listenReturn() {
	ret := <-this.notifyReturn
	if string(ret.Body) != "" {
		log.Println("消息没有正确入列:", string(ret.Body))
	} else {
		log.Println("消息正确入列:")
	}
}

func (this *MQ) Send() (err error) {
	// mandatory:
	//如果为true,在exchange正常且可到达的情况下。如果exchange+routeKey无法投递给queue，那么MQ会将消息返还给生产者
	//如果为false时，则直接丢弃
	fmt.Println(this.Message)
	err = this.Channel.Publish(this.Exchange.Name, this.Key, true, false,
		this.Message,
	)
	if this.Confirm {
		this.listenConfirm()
	}

	return
}
