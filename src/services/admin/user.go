package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/admin"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	BlockUser(userID uint) error
	UnblockUser(userID uint) error
	DeleteUser(userID uint) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) BlockUser(userID uint) error {
	return u.repo.UpdateUserStatus(userID, "blocked")
}

func (u *userUsecase) UnblockUser(userID uint) error {
	return u.repo.UpdateUserStatus(userID, "active")
}

func (u *userUsecase) DeleteUser(userID uint) error {
	return u.repo.DeleteUser(userID)
}
