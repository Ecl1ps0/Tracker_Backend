package DTO

import "time"

type TaskDTO struct {
	Title       string
	Description string
	AccessFrom  *time.Time
	AccessTo    *time.Time
}
