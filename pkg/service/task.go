package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (uint, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id uint) (models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) GetAllTeacherTasks(id uint) ([]models.Task, error) {
	return s.repo.GetAllTeacherTasks(id)
}

func (s *TaskService) GetAllStudentTasks(id uint) ([]models.Task, error) {
	return s.repo.GetAllStudentTasks(id)
}

func (s *TaskService) UpdateTask(task models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
