package app

import (
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/config"
	v1 "github.com/amiosamu/adv-backend-trainee-assignment/internal/controller/http/v1"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/postgres"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/validator"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(path string) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.Fatalf("error reading config: %w", err)
	}
	SetLogger(cfg.Log.Level)
	log.Info("Setting up postgres...")

	pg, err := postgres.New(cfg.PG.URI, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %w", err))
	}

	defer pg.Close()

	log.Info("Initializing repositories...")

	repo := repo.NewRepos(pg)

	log.Info("Initializing service dependencies...")

	dependencies := service.Dependencies{

		Repos: repo,
	}

	services := service.NewService(dependencies)

	log.Info("Initializing handlers and routes...")
	handler := gin.New()
	handler.Validator() = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
