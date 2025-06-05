package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/fathimasithara01/ecommerce/src/services/admin"
)

type DashboardHandler struct {
	service services.DashboardService
}

func NewDashboardHandler(s services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	data, err := h.service.GetDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load dashboard data"})
		return
	}
	c.JSON(http.StatusOK, data)
}
