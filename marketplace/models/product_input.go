package models

type ProductInput struct {
	Name        string  `form:"name" binding:"required"`
	Description string  `form:"description"`
	Price       float32 `form:"price" binding:"required"`
	Category    string  `form:"category" binding:"required"`
}
