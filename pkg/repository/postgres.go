package repository

import (
	"Proctor/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.UserRole{},
		&models.Task{},
		&models.Report{},
		&models.StudentTask{},
		&models.StudentSolution{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
