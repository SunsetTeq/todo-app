package main

import (
	"github.com/SunsetTeq/firstRestApiProject"
	"github.com/SunsetTeq/firstRestApiProject/package/handler"
	"github.com/SunsetTeq/firstRestApiProject/package/repository"
	"github.com/SunsetTeq/firstRestApiProject/package/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(firstRestApiProject.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
