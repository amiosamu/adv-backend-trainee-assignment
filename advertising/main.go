package main

import (
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/pkg/handler"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/pkg/repository"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	DB, err := InitDB()
	if err != nil {
		fmt.Errorf("error connecting to the database: %w", err)
	}
	repo := repository.NewRepository(DB.DB)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			fmt.Errorf("error occured while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Errorf("err occured on server shuttding down: %s", err.Error())
	if err := DB.DB.Close(); err != nil {
		fmt.Errorf("error occured on db connection close: %s", err.Error())
	}
}
