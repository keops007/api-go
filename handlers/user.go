package handlers

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
    {
		ID: 1,
		Email:    "test@test.com",
		Password: "password",
},
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}
