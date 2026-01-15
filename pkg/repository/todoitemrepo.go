package repository

import (
	"TODOAPP/pkg/models"
	"fmt"
)

type TodoListRepo struct {
	data *models.DataStructures
}

func (t *TodoListRepo) Create(list models.Todo) (int, error) {
	list.Id = len(t.data.Todo) + 1 // добавить uuid
	t.data.Todo = append(t.data.Todo, list)
	return list.Id, nil
}

func (t *TodoListRepo) GetAll() ([]models.Todo, error) {
	return t.data.Todo, nil // возвращаем копию
}

func (t *TodoListRepo) GetById(id int) (models.Todo, error) {
	for _, t := range t.data.Todo {
		if t.Id == id {
			return t, nil
		}
	}
	return models.Todo{}, fmt.Errorf("Todo list not found") // используем функцию из utils > slice_utils > GetByFieldsSlice
}

func (r *TodoListRepo) Update(id int, input models.Todo) error {
	// используем функцию из utils > slice_utils > UpdateItemSlice
	for i, t := range r.data.Todo {
		if t.Id == id {
			input.Id = t.Id
			r.data.Todo[i] = input
			return nil
		}
	}
	return fmt.Errorf("Todo list not found")
}

func (r *TodoListRepo) Delete(id int) error {
	// используем функцию из utils > slice_utils > FilteredSlice
	for i, t := range r.data.Todo {
		if t.Id == id {
			r.data.Todo = append(r.data.Todo[:i], r.data.Todo[i+1:]...)
		}
	}
	return fmt.Errorf("Todo list not found")
}
