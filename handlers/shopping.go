package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-api/services"
)

type ShoppingHandler struct {
	shoppingService services.ShoppingService
}

func NewShoppingHandler(shoppingService services.ShoppingService) *ShoppingHandler {
	return &ShoppingHandler{shoppingService: shoppingService}
}

func (h *ShoppingHandler) GetItems(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	items, err := h.shoppingService.GetItems(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eroare la incarcare"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ShoppingHandler) AddItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "numele este obligatoriu"})
		return
	}

	item, err := h.shoppingService.AddItem(userID, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eroare la adaugare"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ShoppingHandler) ToggleItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	item, err := h.shoppingService.ToggleItem(c.Param("id"), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item negasit"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ShoppingHandler) DeleteItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	if err := h.shoppingService.DeleteItem(c.Param("id"), userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item negasit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sters"})
}
