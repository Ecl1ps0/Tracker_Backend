package models

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Logs      []byte `binding:"required"`
	Report    []byte `binding:"required"`
	Clipboard []byte `binding:"required"`
}
