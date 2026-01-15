package service

import (
	"TODOAPP/pkg/models"
	"TODOAPP/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(listId int, item models.TodoItem) (int, error) {
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(listId int) ([]models.TodoItem, error) {
	return s.repo.GetAll(listId)
}

func (s *TodoItemService) GetById(itemId int) (models.TodoItem, error) {
	return s.repo.GetById(itemId)
}

func (s *TodoItemService) Update(itemId int, input models.TodoItem) error {
	return s.repo.Update(itemId, input)
}

func (s *TodoItemService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}
