package rabbitmq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

var channel *amqp.Channel

func Setup(url string) (*amqp.Connection, *amqp.Channel, error) {
	// conn, err := amqp.Dial(url)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// ch, err := conn.Channel()
	// if err != nil {
	// 	return nil, nil, err
	// }
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				log.Println("✅ Connected to RabbitMQ")
				_, err = ch.QueueDeclare("emails", true, false, false, false, nil)
				channel = ch
				return conn, ch, nil
			}
		}

		log.Printf("❌ Failed to connect to RabbitMQ (attempt %d): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	return nil, nil, err
}
