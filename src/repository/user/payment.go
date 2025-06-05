package repository

import (
	"strings"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func IsOrderPaid(orderID uint) (bool, error) {
	var order models.Order
	err := database.DB.Select("status").First(&order, orderID).Error
	if err != nil {
		return false, err
	}
	return strings.ToLower(order.Status) == "paid", nil
}

func SavePayment(payment *models.Payment) error {
	return database.DB.Create(payment).Error
}

func UpdateOrderStatus(orderID uint, status string) error {
	return database.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
