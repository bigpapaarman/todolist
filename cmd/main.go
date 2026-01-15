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
	} // сделай заполненными в отдельном файле
}
func main() {
	ds := initDs()

	//добавить конфиги .yaml

	repos := repository.NewRepository(ds)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(TODOAPP.Server) //TODO меняй NewServer()
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Произошла: %s", err.Error())
	}
}
