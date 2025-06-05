package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func AddToCart(userID uint, req models.AddToCartRequest) error {
	return repository.AddToCart(userID, req)
}

func GetAllCartProducts(userID uint) ([]models.Cart, error) {
	return repository.GetAllCartProducts(userID)
}

func UpdateCartItem(userID, cartID uint, req models.UpdateCartRequest) error {
	return repository.UpdateCartItem(userID, cartID, req)
}

func RemoveCartItem(userID, cartID uint) error {
	return repository.RemoveCartItem(userID, cartID)
}

func ClearCart(userID uint) error {
	return repository.ClearCart(userID)
}
