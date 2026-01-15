package main

import (
	"TODOAPP"
	"TODOAPP/pkg/handler"
	"TODOAPP/pkg/models"
	"TODOAPP/pkg/repository"
	"TODOAPP/pkg/service"
	"log"
)

func initDs() *models.DataStructures {
	return &models.DataStructures{
		Users: []models.User{},
		Todo:  []models.Todo{},
	}
}
func main() {
	ds := initDs()

	repos := repository.NewRepository(ds)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(TODOAPP.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Произошла: %s", err.Error())
	}
}
