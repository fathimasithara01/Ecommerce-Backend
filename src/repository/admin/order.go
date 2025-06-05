package repository

import (
	"github.com/fathimasithara01/ecommerce/utils/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrderStatus(id uint, status string) error
	DeleteOrder(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("User").Preload("OrderItems.Product").Preload("Payment").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("User").Preload("OrderItems.Product").Preload("Payment").
		First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) UpdateOrderStatus(id uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *orderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
