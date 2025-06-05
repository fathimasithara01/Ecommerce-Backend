package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/user"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	SearchProducts(query string) ([]models.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (u *productService) GetAllProducts() ([]models.Product, error) {
	return u.productRepo.GetAllProducts()
}

func (u *productService) GetProductByID(id uint) (*models.Product, error) {
	return u.productRepo.GetProductByID(id)
}

func (u *productService) SearchProducts(query string) ([]models.Product, error) {
	return u.productRepo.SearchProducts(query)
}

// func CreateProduct(data models.Product) (*models.Product, error) {
// 	return repository.CreateProduct(data)
// }

// func GetAllProduct() ([]models.Product, error) {
// 	return repository.GetAllProduct()
// }

// func UpdateProduct(id uint, update models.ProductUpdate) (*models.Product, error) {
// 	return repository.UpdateProduct(id, update)
// }

// func DeleteProduct(id uint) error {
// 	return repository.DeleteProduct(id)
// }

// func SearchProduct(query string) ([]models.Product, error) {
// 	return repository.SearchProduct(query)
// }
