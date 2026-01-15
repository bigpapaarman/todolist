package service

import (
	"TODOAPP/pkg/models"
	"TODOAPP/pkg/repository"
)

type Autorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(list models.Todo) (int, error)
	GetAll() ([]models.Todo, error)
	GetById(id int) (models.Todo, error)
	Update(id int, input models.Todo) (int, error)
	Delete(id int) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(listId int) ([]models.TodoItem, error)
	GetById(itemId int) (models.TodoItem, error)
	Update(itemId int, input models.TodoItem) error
	Delete(itemId int) error
}

type Service struct {
	Autorization Autorization
	TodoList     TodoList
	TodoItem     TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList:     NewTodoListService(repos.TodoList),
		TodoItem:     NewTodoItemService(repos.TodoItem),
	}
}
