package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type RoleService struct {
	repo repository.Role
}

func NewRoleService(repo repository.Role) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) SetDefaultRoles() ([]models.UserRole, error) {
	roles, count, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if count != 0 {
		return roles, nil
	}

	roles, err = s.repo.CreateDefault()
	if err != nil {
		return nil, err
	}

	return roles, nil
}
