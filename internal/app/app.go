package app

import (
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/config"
	_ "github.com/amiosamu/adv-backend-trainee-assignment/docs"
	v1 "github.com/amiosamu/adv-backend-trainee-assignment/internal/controller/http/v1"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/httpserver"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/postgres"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title Advertisement Management Service
// @version 1.0

// @description Test task from avito.tech for a Backend developer trainee.

// @host localhost:8080
// @BasePath  /

func Run(path string) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.Fatalf("error reading config: %w", err)
	}
	SetLogger(cfg.Log.Level)
	log.Info("Setting up postgres...")
	db, err := postgres.InitDB()
	if err != nil {
		log.Fatalf("unable to init database: %v\n", err)
	}

	defer db.Close()

	log.Info("Initializing repositories...")

	repository := repo.NewRepos(db.DB)

	log.Info("Initializing service dependencies...")

	dependencies := service.Dependencies{

		Repos: repository,
	}

	services := service.NewServices(dependencies)

	log.Info("Initializing handlers and routes...")
	handler := gin.New()

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
