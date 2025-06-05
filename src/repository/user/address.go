package repository

import (
	"errors"
	"fmt"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func CreateAddress(userID uint, address *models.Address) error {
	var existAdddress models.Address
	if err := database.DB.Where("user_id = ?", userID).Find(&existAdddress).Error; err != nil {
		return fmt.Errorf("someone used this addrress invlaidd address id")
	}
	return database.DB.Create(&address).Error
}

func GetAllAddress(userID uint) ([]models.Address, error) {
	var address []models.Address
	err := database.DB.Where("user_id =?", userID).Find(&address).Error
	return address, err
}

func UpdateAddress(userID, addressID uint, updated *models.Address) error {
	result := database.DB.Model(&models.Address{}).
		Where("user_id = ? AND id = ?", userID, addressID).
		Updates(updated)

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated: address not found or no changes")
	}

	return result.Error
}

func DeleteAddress(userID, addressID uint) error {
	var address models.Address
	result := database.DB.
		Where("user_id = ? AND id = ?", userID, addressID).
		Delete(&address)

	if result.RowsAffected == 0 {
		return errors.New("address not found or already deleted")
	}
	return result.Error
}
