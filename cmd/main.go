package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexander256/shop/config"
	"github.com/alexander256/shop/logger"
	"github.com/alexander256/shop/pkg/handler"
	"github.com/alexander256/shop/pkg/service"
	"github.com/alexander256/shop/pkg/storage"
	"github.com/alexander256/shop/server"
)

// @title Shop APP API
// @version 1.0
// @description API Server for Shop app Application

// @host localhost:8181
// @BasePath /
func main() {
	config, err := config.InitCinfig()
	if err != nil {
		log.Fatal("error initializing configs :", err)
	}

	log := logger.InitLogger(config.Logger.Level)

	db, err := storage.NewPostgresDB(&config.Postgres)
	if err != nil {
		log.Fatal("failed to initialize db : ", err)
	}

	storage := storage.NewStorage(db, log)
	service := service.NewService(storage, log)
	handler := handler.NewHandler(service, log)

	srv := server.NewServer()

	go func() {
		if err := srv.Run(config.Gin.Port, handler.InitRoutes(config.Gin.Mode)); err != nil {
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
