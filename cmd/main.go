package main

import (
	"github.com/SunsetTeq/firstRestApiProject"
	"github.com/SunsetTeq/firstRestApiProject/package/handler"
	"github.com/SunsetTeq/firstRestApiProject/package/repository"
	"github.com/SunsetTeq/firstRestApiProject/package/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initalizing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(firstRestApiProject.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs") // добавляем директорию конфига
	viper.SetConfigName("config")  // имя конфига
	return viper.ReadInConfig()
}
