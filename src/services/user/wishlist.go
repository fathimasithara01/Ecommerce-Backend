package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func AddToWishlist(userID uint, req models.AddToWishlistRequest) error {
	return repository.AddToWishlist(userID, req)
}

func GetWishlist(userID uint) ([]models.Wishlist, error) {
	return repository.GetWishlist(userID)
}

func RemoveFromWishlist(userID, productID uint) error {
	return repository.RemoveFromWishlist(userID, productID)
}

func ClearWishlist(userID uint) error {
	return repository.ClearWishlist(userID)
}
