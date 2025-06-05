package repository

import (
	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func GetProductByCategoryID(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}
