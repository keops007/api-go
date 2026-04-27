package repository

import (
	"go-api/models"

	"gorm.io/gorm"
)

type ShoppingRepository interface {
	FindByUserID(userID uint) ([]models.ShoppingItem, error)
	FindByIDAndUserID(id string, userID uint) (*models.ShoppingItem, error)
	Create(item *models.ShoppingItem) error
	SetDone(id uint, done bool) error
	Delete(id string, userID uint) (int64, error)
}

type shoppingRepository struct {
	db *gorm.DB
}

func NewShoppingRepository(db *gorm.DB) ShoppingRepository {
	return &shoppingRepository{db: db}
}

func (r *shoppingRepository) FindByUserID(userID uint) ([]models.ShoppingItem, error) {
	var items []models.ShoppingItem
	err := r.db.Where("user_id = ?", userID).Order("done asc, created_at desc").Find(&items).Error
	return items, err
}

func (r *shoppingRepository) FindByIDAndUserID(id string, userID uint) (*models.ShoppingItem, error) {
	var item models.ShoppingItem
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *shoppingRepository) Create(item *models.ShoppingItem) error {
	return r.db.Create(item).Error
}

func (r *shoppingRepository) SetDone(id uint, done bool) error {
	return r.db.Model(&models.ShoppingItem{}).Where("id = ?", id).Update("done", done).Error
}

func (r *shoppingRepository) Delete(id string, userID uint) (int64, error) {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.ShoppingItem{})
	return result.RowsAffected, result.Error
}
