package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	services "github.com/fathimasithara01/ecommerce/src/services/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
	"github.com/fathimasithara01/ecommerce/utils/response"
)

func CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}

	if err := validator.New().Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation falied", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.CreateOrder(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not create order", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "order created", nil, nil))
}

func GetUserOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	orders, err := services.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not fetch orders", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "orders fetched", orders, nil))
}

func GetOrderById(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	if orderIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"message":     "order ID is required",
			"data":        nil,
			"error":       "missing order_id in URL",
		})
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"message":     "invalid order id",
			"data":        nil,
			"error":       err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")
	order, err := services.GetOrderById(uint(orderID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, "order not found", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "order fetched", order, nil))
}
