package repository

import (
	"Proctor/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	userTable = "users"
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
		&models.Data{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
