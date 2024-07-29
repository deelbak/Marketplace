package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"marketPlace/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port = 5432 user=deelbak dbname=deelbak password=passw0rd sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Product{}, &models.User{})
	DB = db
	return db, nil
}
