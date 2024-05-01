package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(userId uint) (models.User, error) {
	user, err := s.repo.GetProfile(userId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
