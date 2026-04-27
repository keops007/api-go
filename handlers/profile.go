package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go-api/services"
)

type ProfileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

var allowedExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true,
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	user, err := h.profileService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     user.ID,
		"email":  user.Email,
		"avatar": user.Avatar,
	})
}

func (h *ProfileHandler) UploadAvatar(c *gin.Context) {
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
	if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eroare la salvare"})
		return
	}

	avatarURL := "/uploads/" + filename
	if err := h.profileService.UpdateAvatar(userID, avatarURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eroare la actualizare"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"avatar": avatarURL})
}
