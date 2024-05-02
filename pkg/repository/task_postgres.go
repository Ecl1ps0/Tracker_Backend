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

func (r *TaskPostgres) GetTaskByID(id uint) (models.Task, error) {
	var task models.Task
	if result := r.db.Where("id = ?", id).First(&task); result.Error != nil {
		fmt.Println(result)
		return models.Task{}, result.Error
	}

	return task, nil
}

func (r *TaskPostgres) DeleteTask(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
