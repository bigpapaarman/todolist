package service

import (
	"TODOAPP/pkg/models"
	"TODOAPP/pkg/repository"
)

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user := models.User{
		Username: username,
		Password: password,
	}

	token, err := s.repo.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	parsedToken, err := s.repo.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return parsedToken, nil
}
