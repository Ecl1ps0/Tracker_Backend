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

func (s *UserService) AddStudentToTask(studentId, taskId uint) error {
	return s.repo.AddStudentToTask(models.StudentTask{
		StudentID: studentId,
		TaskID:    taskId,
	})
}

func (s *UserService) GetRoleByUserID(id uint) (uint, error) {
	return s.repo.GetRoleByID(id)
}
