package mail

import (
	"log"
	"mailqueue/internal/models"
)

func Send(email models.Email) {
	// Simulate email sending
	log.Printf("[MailQueue] Sending email %s with subject %s", email.To, email.Subject)
}
