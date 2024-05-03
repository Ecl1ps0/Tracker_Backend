package repository

import (
	"Proctor/models"
	"fmt"
	"gorm.io/gorm"
)

type TaskPostgres struct {
	db *gorm.DB
}

func NewTaskPostgres(db *gorm.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) CreateTask(task models.Task) (uint, error) {
	if result := r.db.Create(&task); result.Error != nil {
		return 0, result.Error
	}

	return task.ID, nil
}

func (r *TaskPostgres) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if result := r.db.Find(&tasks); result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskPostgres) GetTaskByID(id uint) (models.Task, error) {
	var task models.Task
	if result := r.db.Where("id = ?", id).First(&task); result.Error != nil {
		fmt.Println(result)
		return models.Task{}, result.Error
	}

	return task, nil
}

func (r *TaskPostgres) GetAllTeacherTasks(id uint) ([]models.Task, error) {
	var tasks []models.Task
	if result := r.db.Where("teacher_id = ?", id).Find(&tasks); result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskPostgres) GetAllStudentTasks(id uint) ([]models.Task, error) {
	var tasks []models.Task
	if result := r.db.Joins("join student_tasks on tasks.id = student_tasks.task_id").Where("student_tasks.student_id = ?", id).Find(&tasks); result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskPostgres) UpdateTask(task models.Task) error {
	if result := r.db.Save(&task); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskPostgres) DeleteTask(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
