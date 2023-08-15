package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	wbl0 "github.com/pamallika/WBL0"
	"github.com/pamallika/WBL0/pkg/handler"
	"github.com/pamallika/WBL0/pkg/repository"
	"github.com/pamallika/WBL0/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("config connection error: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("db connection error: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(wbl0.Server)
	err = srv.Run("8000", handlers.InitRouts())
	if err != nil {
		log.Fatalf("server init error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
