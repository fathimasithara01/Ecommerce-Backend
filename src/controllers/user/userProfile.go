package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/fathimasithara01/ecommerce/src/services/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
	"github.com/fathimasithara01/ecommerce/utils/response"
)

func CreateUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "User not authenticated", nil, nil))
		return
	}

	var req models.CreateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}

	err := services.CreateUserProfile(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "couldn't created user profile", nil, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "profile created successfully ", nil, nil))

}

func GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "User not authenticated", nil, nil))
		return
	}

	profile, err := services.GetUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "there is no userProfile", nil, err.Error()))
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message":     "User profile not found",
			"error":       "profile not found",
		})
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "profile fetched successfully", profile, nil))

}

func UpdateUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "User not authenticated", nil, nil))
		return
	}

	var req models.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}

	err := services.UpdateUserProfile(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not get cart", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "profile updated successfully", nil, nil))

}

func DeleteUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "User not authenticated", nil, nil))
		return
	}

	err := services.DeleteUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not get cart", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "profile deleted successfully", nil, nil))

}
