package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/im-faix/sentinel/backend/internal/config"
	"github.com/im-faix/sentinel/backend/internal/logger"
	"github.com/im-faix/sentinel/backend/internal/server"
)

type Application struct {
	Config *config.Config
	Logger *slog.Logger
	Server *server.Server
	Router *chi.Mux
}

func New() *Application {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel)

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Sentinel API"))
	})

	srv := server.New(cfg.Address(), router, log)

	return &Application{
		Config: cfg,
		Logger: log,
		Server: srv,
		Router: router,
	}
}

func (a *Application) Run() error {

	go func() {
		if err := a.Server.Start(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {

			a.Logger.Error(
				"server failed",
				slog.String("error", err.Error()),
			)

			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		return err
	}

	a.Logger.Info("Sentinel stopped")

	return nil
}
