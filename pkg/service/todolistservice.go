package service

import (
	"TODOAPP/pkg/models"
	"TODOAPP/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(list models.Todo) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoListService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoListService) GetById(id int) (models.Todo, error) {
	return s.repo.GetById(id)
}

func (s *TodoListService) Update(id int, input models.Todo) (int, error) {
	err := s.repo.Update(id, input)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TodoListService) Delete(id int) error {
	return s.repo.Delete(id)
}
