package main

import (
	"log"
	"net/http"

	"github.com/budsx/project-layout/config"
	"github.com/budsx/project-layout/pkg/httpserver"
	logs "github.com/budsx/project-layout/pkg/logger"
	"github.com/budsx/project-layout/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load Configuration
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logging
	logger := logs.NewLogger(conf.Log.Level)

	// Postgres
	// pg, err := postgres.NewPostgres(conf.PG.URL, postgres.MaxPoolSize(conf.PG.PoolMax))
	// if err != nil {
	// 	logger.Fatal(fmt.Errorf("postgres.NewPostgres: %w", err))
	// }
	// defer pg.Close()

	// Usecase

	// HTTP Server
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.RequestURI)
		w.WriteHeader(http.StatusOK)
	})
	httpServer := httpserver.NewHTTPServer(r, httpserver.Port(conf.HTTP.Port))

	// Shutdown
	utils.OnShutdown(func() {
		httpServer.Shutdown()
	})
}
