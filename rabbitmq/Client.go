package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Client struct {
	Conn *amqp.Connection
}

func New() *Client {
	client := &Client{}
	client.Connection()
	return client
}

func (this *Client)Connection()  {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "admin", "admin", "192.168.8.99", 5672)
	conn, err := amqp.Dial(dsn)
	if err!= nil {
		log.Fatal(err)
	}
	this.Conn = conn
}

func (this *Client)Close()  {
	this.Conn.Close()
}
