package repository

import (
	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/utils/models"
)

type DashboardRepository interface {
	CountUsers() (int64, error)
	CountOrders() (int64, error)
	SumRevenue() (float64, error)
	CountOrdersByStatus() (map[string]int64, error)
}

type dashboardRepo struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepo{db}
}

func (r *dashboardRepo) CountUsers() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return count, err
}

func (r *dashboardRepo) CountOrders() (int64, error) {
	var count int64
	err := r.db.Model(&models.Order{}).Count(&count).Error
	return count, err
}

func (r *dashboardRepo) SumRevenue() (float64, error) {
	var total *float64
	err := r.db.Model(&models.Order{}).
		Where("status IN ?", []string{"delivered", "shipped"}).
		Select("SUM(total_amount)").Scan(&total).Error
	if err != nil {
		return 0, err
	}
	if total == nil {
		return 0, nil // no rows matched, sum is zero
	}
	return *total, nil
}

func (r *dashboardRepo) CountOrdersByStatus() (map[string]int64, error) {
	var results []struct {
		Status string
		Count  int64
	}
	err := r.db.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&results).Error

	statusMap := make(map[string]int64)
	for _, r := range results {
		statusMap[r.Status] = r.Count
	}
	return statusMap, err
}
