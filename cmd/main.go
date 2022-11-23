package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexander256/shop/models"
	"github.com/alexander256/shop/pkg/handler"
	"github.com/alexander256/shop/pkg/service"
	"github.com/alexander256/shop/pkg/storage"
	"github.com/alexander256/shop/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	cfgPostgres := models.ConfigPostgres{
		Host:     "localhost",
		Port:     "5432",
		UserName: "postgres",
		Password: "postgres",
		DBName:   "shop",
		SSLMode:  "disable",
	}

	db, err := storage.NewPostgresDB(&cfgPostgres)
	if err != nil {
		log.Fatal("failed to initialize db : ", err)
	}

	storage := storage.NewStorage(db, log)
	service := service.NewService(storage, log)
	handler := handler.NewHandler(service, log)

	srv := server.NewServer()

	go func() {
		if err := srv.Run("8080", handler.InitRoutes("debug")); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Info("app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("app shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
