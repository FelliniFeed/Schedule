package main

import (
	"context"
	"os"
	"os/signal"
	"schedule/pkg"
	"schedule/pkg/handler"
	"schedule/pkg/repository"
	"schedule/pkg/service"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {

	//TODO: Сделать миграцию БД
	// Дописать все API
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occired initializing configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error occired initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("database.dbname"),
		SSLMode: viper.GetString("database.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(pkg.Server)

	go func () {
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occired while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<- quit

	logrus.Print("Server stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occired on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occired on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}