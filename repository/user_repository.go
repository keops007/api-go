package repository

import (
	"go-api/models"

	"gorm.io/gorm"
)

// UserRepository defineste contractul pentru accesul la date despre useri.
// Orice implementare (DB reala, mock pentru teste) trebuie sa respecte aceasta interfata.
type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) error
	UpdateAvatar(id uint, avatarURL string) error
}

// userRepository este implementarea concreta care foloseste GORM + PostgreSQL.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returneaza o instanta noua. Returnam interfata, nu struct-ul concret —
// astfel codul care apeleaza aceasta functie nu stie nimic despre GORM.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateAvatar(id uint, avatarURL string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("avatar", avatarURL).Error
}
