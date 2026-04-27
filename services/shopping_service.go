package services

import (
	"errors"

	"go-api/models"
	"go-api/repository"
)

type ShoppingService interface {
	GetItems(userID uint) ([]models.ShoppingItem, error)
	AddItem(userID uint, name string) (*models.ShoppingItem, error)
	ToggleItem(id string, userID uint) (*models.ShoppingItem, error)
	DeleteItem(id string, userID uint) error
}

type shoppingService struct {
	repo repository.ShoppingRepository
}

func NewShoppingService(repo repository.ShoppingRepository) ShoppingService {
	return &shoppingService{repo: repo}
}

func (s *shoppingService) GetItems(userID uint) ([]models.ShoppingItem, error) {
	return s.repo.FindByUserID(userID)
}

func (s *shoppingService) AddItem(userID uint, name string) (*models.ShoppingItem, error) {
	item := &models.ShoppingItem{UserID: userID, Name: name}
	err := s.repo.Create(item)
	return item, err
}

func (s *shoppingService) ToggleItem(id string, userID uint) (*models.ShoppingItem, error) {
	item, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, errors.New("item not found")
	}
	item.Done = !item.Done
	err = s.repo.SetDone(item.ID, item.Done)
	return item, err
}

func (s *shoppingService) DeleteItem(id string, userID uint) error {
	rows, err := s.repo.Delete(id, userID)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("item not found")
	}
	return nil
}
