package DTO

type UserDTO struct {
	ID      uint
	Name    string
	Surname *string
	Email   string
	RoleID  uint
}
