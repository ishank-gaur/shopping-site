package handlers

import (
	"ecommerce/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func PlaceOrder(c *gin.Context) {
	userID := c.GetString("user")
	var cart models.Cart
	if err := models.DB.Preload("Items").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	order := models.Order{UserID: cart.UserID, Items: cart.Items}
	models.DB.Create(&order)
	models.DB.Delete(&cart)
	c.JSON(http.StatusOK, order)
}

func ListOrders(c *gin.Context) {
	userID := c.GetString("user")
	var orders []models.Order
	models.DB.Preload("Items").Where("user_id = ?", userID).Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func toUint(s string) uint {
	var i uint
	fmt.Sscanf(s, "%d", &i)
	return i
}