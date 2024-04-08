package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type FileHandlerPostgres struct {
	db *gorm.DB
}

func NewFileHandlerPostgres(db *gorm.DB) *FileHandlerPostgres {
	return &FileHandlerPostgres{db: db}
}

func (r *FileHandlerPostgres) Create(data models.Data) (uint, error) {
	if result := r.db.Create(&data); result.Error != nil {
		return 0, nil
	}

	return data.ID, nil
}
