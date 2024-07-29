package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"marketPlace/handlers"
	"marketPlace/utils"
	"net/http"
)

func main() {
	r := gin.Default()

	_, err := utils.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	r.LoadHTMLGlob("templates/*")
	r.GET("/create-user", func(c *gin.Context) {
		log.Println("GET /create-user called")
		c.HTML(http.StatusOK, "create_user.html", nil)
	})
	r.GET("/users", handlers.GetAllUsers)

	r.POST("/create-user", handlers.CreateUser)
	log.Println("Starting server on :8080")
	r.Run(":8080")
	//fmt.Println("Hello World")
}
