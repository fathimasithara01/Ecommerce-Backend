package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	services "github.com/fathimasithara01/ecommerce/src/services/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
	"github.com/fathimasithara01/ecommerce/utils/response"
)

func MakePayment(c *gin.Context) {
	var req models.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}
	if err := validator.New().Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.ProcessPayment(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "payment failed", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "payment proceed", nil, nil))
}
