package models

type UserInput struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
