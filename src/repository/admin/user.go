// internal/repository/user_repository.go
package repository

import (
	"errors"

	"github.com/fathimasithara01/ecommerce/utils/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	UpdateUserStatus(userID uint, status string) error
	DeleteUser(userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	// if err := r.db.Preload("UserProfile").Find(&users).Error; err != nil {
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUserStatus(userID uint, status string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("status", status).Error
}

func (r *userRepository) DeleteUser(userID uint) error {
	res := r.db.Delete(&models.User{}, userID)
	if res.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return res.Error
}
