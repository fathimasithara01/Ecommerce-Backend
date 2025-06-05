package repository

import (
	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/utils/models"
)

type AdminRepository struct {
	DB *gorm.DB
}

func (r *AdminRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *AdminRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}
