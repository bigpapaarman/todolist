package repository

import (
	"TODOAPP/pkg/models"
	"errors"
	"fmt"
)

type TodoItemRepo struct {
	data *models.DataStructures
}

func (r *TodoItemRepo) Create(listId int, item models.TodoItem) (int, error) {
	item.ListID = listId
	item.Id = len(r.data.Item) + 1
	r.data.Item = append(r.data.Item, item)
	return item.Id, nil
}

func (r *TodoItemRepo) GetAll(listId int) ([]models.TodoItem, error) {
	var result []models.TodoItem
	for _, item := range r.data.Item {
		if item.ListID == listId {
			result = append(result, item)
		}
	}
	return result, nil
}

func (r *TodoItemRepo) GetById(itemId int) (models.TodoItem, error) {
	for _, item := range r.data.Item {
		if item.Id == itemId {
			return item, nil
		}
	}
	return models.TodoItem{}, errors.New("item not found")
}

func (r *TodoItemRepo) Update(itemId int, input models.TodoItem) error {
	for i, item := range r.data.Item {
		if item.Id == itemId {
			input.Id = item.Id
			input.ListID = item.ListID
			r.data.Item[i] = input
			return nil

		}
	}
	return fmt.Errorf("item not found")
}

func (r *TodoItemRepo) Delete(itemId int) error {
	for i, item := range r.data.Item {
		if item.Id == itemId {
			r.data.Item = append(r.data.Item[:i], r.data.Item[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}
