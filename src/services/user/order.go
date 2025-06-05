package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func CreateOrder(userID uint, req models.CreateOrderRequest) error {
	var total float64
	var items []models.OrderItem

	for _, item := range req.Items {
		product, err := repository.GetProductById(item.ProductID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("product with ID %d not found", item.ProductID)
			}
			return fmt.Errorf("error fetching product: %v", err)
		}

		price := product.Price * float64(item.Quantity)
		total += price

		items = append(items, models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     price,
		})
	}

	order := models.Order{
		UserID: userID,
		// TotalAmount: total,
		AddressID:   req.AddressID,
		TotalAmount: total,
		Status:      "pending",
		// ShippingAddress: ,
		OrderItems: items,
	}

	return repository.CreateOrder(&order)
}

func GetUserOrders(userID uint) ([]models.Order, error) {
	return repository.GetUserOrders(userID)
}

func GetOrderById(orderID, userID uint) (models.Order, error) {
	return repository.GetOrderById(orderID, userID)
}
