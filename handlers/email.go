package handlers

import (
	"encoding/json"
	"mailqueue/internal/models"
	"mailqueue/internal/rabbitmq"
	"net/http"
)

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Email
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err := rabbitmq.Publish("emails", req)
	if err != nil {
		http.Error(w, "Failed to enqueue email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Email enqueued",
	})
}
