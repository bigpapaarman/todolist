package repository

import (
	"TODOAPP/pkg/models"
	"errors"
)

type AuthRepo struct {
	data *models.DataStructures
}

func (a *AuthRepo) CreateUser(user models.User) (int, error) {
	user.Id = len(a.data.Users) + 1
	a.data.Users = append(a.data.Users, user)
	return user.Id, nil
}

func (a *AuthRepo) GetByUsername(username string) (models.User, error) {
	for _, u := range a.data.Users {
		if u.Username == username {
			return u, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (a *AuthRepo) GenerateToken(user models.User) (string, error) {
	return "", nil
}

func (a *AuthRepo) ParseToken(token string) (int, error) {
	return 0, nil
}
