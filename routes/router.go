package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	return r
}
