package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-api/config"
	"go-api/models"
)

func GetShoppingItems(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var items []models.ShoppingItem
	config.DB.Where("user_id = ?", userID).Order("done asc, created_at desc").Find(&items)
	c.JSON(http.StatusOK, items)
}

func AddShoppingItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "numele este obligatoriu"})
		return
	}

	item := models.ShoppingItem{UserID: userID, Name: input.Name}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func ToggleShoppingItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var item models.ShoppingItem
	if err := config.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item negasit"})
		return
	}

	config.DB.Model(&item).Update("done", !item.Done)
	item.Done = !item.Done
	c.JSON(http.StatusOK, item)
}

func DeleteShoppingItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	result := config.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).Delete(&models.ShoppingItem{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "item negasit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sters"})
}
