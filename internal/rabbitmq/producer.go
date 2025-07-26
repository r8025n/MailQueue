package rabbitmq

import (
	"github.com/streadway/amqp"
)

func Setup(url string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = ch.QueueDeclare("emails", true, false, false, false, nil)

	return conn, ch, nil
}
