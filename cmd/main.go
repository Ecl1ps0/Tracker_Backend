package main

import (
	"Proctor"
	"Proctor/pkg/handler"
	"Proctor/pkg/repository"
	"Proctor/pkg/service"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

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
	handlers := handler.NewHandler(services)

	server := new(Proctor.Server)
	if err := server.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
