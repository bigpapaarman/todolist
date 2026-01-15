package repository

import (
	"TODOAPP/pkg/models"
)

type Autorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.User) (string, error)
	GetByUsername(username string) (models.User, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(list models.Todo) (int, error)
	GetAll() ([]models.Todo, error)
	GetById(id int) (models.Todo, error)
	Update(id int, input models.Todo) error
	Delete(id int) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(listId int) ([]models.TodoItem, error)
	GetById(itemId int) (models.TodoItem, error)
	Update(itemId int, input models.TodoItem) error
	Delete(itemId int) error
}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(ds *models.DataStructures) *Repository {
	return &Repository{
		Autorization: &AuthRepo{data: ds},
		TodoList:     &TodoListRepo{data: ds},
		TodoItem:     &TodoItemRepo{data: ds},
	}
}
