package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func GetAllCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

func GetProductByCategoryID(categoryID uint) ([]models.Product, error) {
	return repository.GetProductByCategoryID(categoryID)
}
