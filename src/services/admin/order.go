package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/admin"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

type OrderService interface {
	GetAllOrder() ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrderStatus(id uint, status string) error
	DeleteOrder(id uint) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) GetAllOrder() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *orderService) UpdateOrderStatus(id uint, status string) error {
	return s.repo.UpdateOrderStatus(id, status)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.repo.DeleteOrder(id)
}
