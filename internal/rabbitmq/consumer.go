package rabbitmq

import (
	"encoding/json"
	"log"
	"mailqueue/internal/mail"
	"mailqueue/internal/models"

	"github.com/streadway/amqp"
)

func StartConsumer(ch amqp.Channel) {
	msgs, err := ch.Consume("emails", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	for msg := range msgs {
		var email models.Email
		if err := json.Unmarshal(msg.Body, &email); err != nil {
			mail.Send(email)
		}
	}
}
