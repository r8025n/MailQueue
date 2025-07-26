package rabbitmq

import (
	"encoding/json"
	"mailqueue/internal/models"

	"github.com/streadway/amqp"
)

func Publish(queue string, msg models.Email) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
