package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rgraterol/accounts-core-api/application/responses"
)

func Routes(r *chi.Mux) {
	r.Get("/ping", basePingHandler)
}

func basePingHandler(w http.ResponseWriter, _ *http.Request) {
	responses.OK(w, "pong")
}
