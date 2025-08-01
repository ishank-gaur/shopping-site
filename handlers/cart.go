package handlers

import (
	"ecommerce/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	userID := c.GetString("user")
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if err := models.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: toUint(userID)}
		models.DB.Create(&cart)
	}
	models.DB.Model(&cart).Association("Items").Append(&item)
	c.JSON(http.StatusOK, cart)
}

func ListCart(c *gin.Context) {
	userID := c.GetString("user")
	var cart models.Cart
	if err := models.DB.Preload("Items").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, cart)
}