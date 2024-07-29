package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string
	Price       float64 `gorm:"size:255"`
	Category    string
}

func CreateProduct(db *gorm.DB, name string, description string, price float64, category string) (*Product, error) {
	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}
	if err := db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func GetAllProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
