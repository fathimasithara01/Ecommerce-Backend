package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

func CreateAddress(userID uint, req models.CreateAddressRequest) error {

	address := models.Address{
		UserID:   userID,
		FullName: req.FullName,
		Phone:    req.Phone,
		House:    req.House,
		Street:   req.Street,
		City:     req.City,
		State:    req.State,
		Pincode:  req.Pincode,
	}
	return repository.CreateAddress(userID, &address)
}

func GetAllAdddress(userID uint) ([]models.Address, error) {
	return repository.GetAllAddress(userID)
}

func UpdateAddress(userID, addressID uint, req models.UpdateAddressRequest) error {
	address := models.Address{
		FullName: req.FullName,
		Phone:    req.Phone,
		House:    req.House,
		Street:   req.Street,
		City:     req.City,
		State:    req.State,
		Pincode:  req.Pincode,
	}

	return repository.UpdateAddress(userID, addressID, &address)
}

func DeleteAddress(userID, addressID uint) error {
	return repository.DeleteAddress(userID, addressID)
}
