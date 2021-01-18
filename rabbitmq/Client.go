package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var Conn *amqp.Connection
var once sync.Once
var client *rabbitmq

type rabbitmq struct {
	Conn *amqp.Connection
}

func New(attr ...*attribute) *rabbitmq {
	once.Do(func() {
		client = &rabbitmq{}
		host := Attributes(attr).String(HOST)
		if host == "" {
			host = "127.0.0.1"
		}
		port := Attributes(attr).String(POST)
		if port == "" {
			port = "5672"
		}
		username := Attributes(attr).String(USER_NAME)
		if username == "" {
			username = "admin"
		}
		password := Attributes(attr).String(PASSWORD)
		if password == "" {
			password = "admin"
		}
		client.Connection(host, port, username, password)
		Conn = client.Conn
	})
	return client
}

func (this *rabbitmq) Connection(host string, port string, username string, password string) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	this.Conn = conn
}

func (this *rabbitmq) Close() {
	this.Conn.Close()
}

