package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go-api/config"
	"go-api/models"
)

var allowedExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true,
}

func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     user.ID,
		"email":  user.Email,
		"avatar": user.Avatar,
	})
}

func UploadAvatar(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fisier lipsa"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "format neacceptat (jpg, png, webp)"})
		return
	}

	os.MkdirAll("./uploads", 0755)
	filename := fmt.Sprintf("%d%s", userID, ext)
	savePath := "./uploads/" + filename

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eroare la salvare"})
		return
	}

	avatarURL := "/uploads/" + filename
	config.DB.Model(&models.User{}).Where("id = ?", userID).Update("avatar", avatarURL)

	c.JSON(http.StatusOK, gin.H{"avatar": avatarURL})
}
