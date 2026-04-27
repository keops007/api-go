package services

import (
	"go-api/models"
	"go-api/repository"
)

type ProfileService interface {
	GetProfile(userID uint) (*models.User, error)
	UpdateAvatar(userID uint, avatarURL string) error
}

type profileService struct {
	userRepo repository.UserRepository
}

func NewProfileService(userRepo repository.UserRepository) ProfileService {
	return &profileService{userRepo: userRepo}
}

func (s *profileService) GetProfile(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *profileService) UpdateAvatar(userID uint, avatarURL string) error {
	return s.userRepo.UpdateAvatar(userID, avatarURL)
}
