package services

import (
	"errors"

	repository "github.com/fathimasithara01/ecommerce/src/repository/admin"
	"github.com/fathimasithara01/ecommerce/utils/helper"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

type AuthService struct {
	Repo *repository.AdminRepository
}

func (s *AuthService) SignupAdmin(name, email, password string) (*models.User, error) {
	hasPassword, _ := helper.HashPassword(password)
	user := models.User{Name: name, Email: email, Password: hasPassword, IsAdmin: true}
	return &user, s.Repo.CreateUser(&user)
}

func (s *AuthService) LoginAdmin(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil || !user.IsAdmin {
		return "", errors.New("admin not found or not authorized")
	}

	if !helper.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := helper.GenerateAdminJWT(user.ID, user.IsAdmin)
	return token, err
}
