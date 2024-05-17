package main

import (
	"Proctor"
	"Proctor/pkg/handler"
	"Proctor/pkg/repository"
	"Proctor/pkg/service"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Tracker API
// @version 1.0
// @description API Server for Tracker Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Fail to read config file: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Fail to load env: %v", err)
	}

	db, err := repository.NewPostgresDB(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		viper.GetString("db.user"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname")),
	)
	if err != nil {
		logrus.Fatalf("Fail to initialize database connection: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	roles, err := services.Role.SetDefaultRoles()
	if err != nil {
		logrus.Fatalf("Fail to set default roles: %v", err)
	}
	logrus.Infof("Successfully set default roles: %s, %s, %s", roles[0].Name, roles[1].Name, roles[2].Name)

	handlers := handler.NewHandler(services)

	server := new(Proctor.Server)
	go func() {
		if err := server.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %v", err)
		}
	}()

	logrus.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("App shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
