package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	services "github.com/fathimasithara01/ecommerce/src/services/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
	"github.com/fathimasithara01/ecommerce/utils/response"
)

func UserSignUp(c *gin.Context) {
	var req models.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil || validator.New().Struct(req) != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}

	err := services.UserSignUp(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "signup failed", req, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "signup successful, OTP sent", req, nil))

}

func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil || validator.New().Struct(req) != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid login data", nil, nil))
		return
	}

	token, err := services.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "login failed", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "login successful", gin.H{"token": token}, nil))

}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Logout successful",
	})
}
