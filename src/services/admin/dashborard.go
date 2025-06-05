package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/admin"
)

type DashboardData struct {
	TotalUsers     int64            `json:"total_users"`
	TotalOrders    int64            `json:"total_orders"`
	TotalRevenue   float64          `json:"total_revenue"`
	OrdersByStatus map[string]int64 `json:"orders_by_status"`
}

type DashboardService interface {
	GetDashboardStats() (*DashboardData, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(r repository.DashboardRepository) DashboardService {
	return &dashboardService{repo: r}
}

func (s *dashboardService) GetDashboardStats() (*DashboardData, error) {
	users, err := s.repo.CountUsers()
	if err != nil {
		return nil, err
	}

	orders, err := s.repo.CountOrders()
	if err != nil {
		return nil, err
	}

	revenue, err := s.repo.SumRevenue()
	if err != nil {
		return nil, err
	}

	statusCounts, err := s.repo.CountOrdersByStatus()
	if err != nil {
		return nil, err
	}

	return &DashboardData{
		TotalUsers:     users,
		TotalOrders:    orders,
		TotalRevenue:   revenue,
		OrdersByStatus: statusCounts,
	}, nil
}
