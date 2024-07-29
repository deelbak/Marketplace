package handlers

import (
	"github.com/gin-gonic/gin"
	"marketPlace/models"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var productInput models.ProductInput
	if err := c.ShouldBind(&productInput); err != nil {
		c.HTML(http.StatusBadRequest, "create_product.html", gin.H{"error": err.Error()})
		return
	}
}
