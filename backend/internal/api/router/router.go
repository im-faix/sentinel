package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/im-faix/sentinel/backend/internal/health"
	"github.com/im-faix/sentinel/backend/internal/version"
)

func New() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte("Sentinel API"))
	})

	r.Get("/health", health.Health)

	r.Get("/live", health.Live)

	r.Get("/ready", health.Ready)

	r.Route("/api/v1", func(api chi.Router) {

		api.Get("/version", version.Get)

	})

	return r
}
