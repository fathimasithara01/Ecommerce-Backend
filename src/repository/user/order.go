package repository

import (
	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func GetProductById(id uint) (models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func CreateOrder(order *models.Order) error {

	err := database.DB.Create(&order).Error

	return err
}

func GetUserOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Preload("OrderItems.Product").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func GetOrderById(orderID, userID uint) (models.Order, error) {
	var order models.Order
	err := database.DB.Preload("OrderItems.Product").Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error
	return order, err
}
