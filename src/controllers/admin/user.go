package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	services "github.com/fathimasithara01/ecommerce/src/services/admin"
)

type UserHandler struct {
	usecase services.UserUsecase
}

func NewUserHandler(usecase services.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) BlockUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.usecase.BlockUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not block user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user blocked successfully"})
}

func (h *UserHandler) UnblockUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.usecase.UnblockUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not unblock user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user unblocked successfully"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.usecase.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
