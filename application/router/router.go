package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rgraterol/accounts-core-api/application/handlers"
	"github.com/rgraterol/accounts-core-api/application/middlewares"
	"github.com/rgraterol/accounts-core-api/application/repositories"
	"github.com/rgraterol/accounts-core-api/application/responses"
	"github.com/rgraterol/accounts-core-api/domain/services"
)

func Routes(r *chi.Mux) {
	metricsMiddleware := middlewares.NewMetrics()

	r.Use(metricsMiddleware.Log)

	r.Get("/ping", basePingHandler)

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/users", func(r chi.Router) {
		s := buildProductiveUsersInterface()
		r.Post("/", handlers.UsersNews(&s))
		r.Get("/{userID}", handlers.GetAccount(&s))
	})

	r.Route("/movements", func(r chi.Router) {
		s := buildProductiveMovementsInterface()
		r.Post("/deposit/{userID}", handlers.MakeDeposit(&s))
		r.Post("/transfer", handlers.MakeMovement(&s))
	})

}

func basePingHandler(w http.ResponseWriter, _ *http.Request) {
	responses.OK(w, "pong")
}

func buildProductiveUsersInterface() services.Users {
	return services.Users{
		AccountsService: &services.Accounts{
			Repository: &repositories.Accounts{},
		},
	}
}

func buildProductiveMovementsInterface() services.Movements {
	return services.Movements{
		AccountsRepository:  &repositories.Accounts{},
		MovementsRepository: &repositories.Movements{},
	}
}
