package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"marketPlace/models"
	"marketPlace/utils"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userInput models.UserInput
	if err := c.ShouldBind(&userInput); err != nil {
		c.HTML(http.StatusBadRequest, "create_user.html", gin.H{"error": err.Error()})
		return
	}
	db := utils.DB
	//defer db.Close()

	user, err := models.CreateUser(db, userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "create_user.html", gin.H{"error": "Қолданушы туыла алмай қалды"})
		return
	}
	c.HTML(http.StatusOK, "create_user.html", gin.H{"message": "Қолданушы сәтті туылды", "user": user})
}
func GetAllUsers(c *gin.Context) {
	db := utils.DB
	user, err := models.GetAllUsers(db)
	if err != nil {
		log.Printf("Error getting all users: %s\n", err.Error())
		c.HTML(http.StatusInternalServerError, "list_users.html", gin.H{"error": "Users not found"})
		return
	}
	c.HTML(http.StatusOK, "list_users.html", gin.H{"users": user})

}
