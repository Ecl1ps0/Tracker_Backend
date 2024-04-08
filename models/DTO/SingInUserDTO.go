package DTO

type SignInUserDTO struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
