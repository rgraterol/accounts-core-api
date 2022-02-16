package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rgraterol/accounts-core-api/application/handlers"
	"github.com/rgraterol/accounts-core-api/application/responses"
	"github.com/rgraterol/accounts-core-api/domain/services"
)

func Routes(r *chi.Mux) {
	r.Get("/ping", basePingHandler)

	r.Route("/users", func(r chi.Router) {
		s := buildProductiveUsersInterface()
		r.Post("/", handlers.UsersNews(&s))
	})
}

func basePingHandler(w http.ResponseWriter, _ *http.Request) {
	responses.OK(w, "pong")
}

func buildProductiveUsersInterface() services.Users {
	return services.Users{
		AccountsService: &services.Accounts{},
	}
}
