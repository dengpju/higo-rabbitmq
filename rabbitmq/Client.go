package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func Connection()  {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "admin", "admin", "192.168.8.99", 5672)
	conn, err := amqp.Dial(dsn)
	if err!= nil {
		log.Fatal(err)
	}
	defer conn.Close()

}
