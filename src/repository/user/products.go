package repository

import (
	"github.com/fathimasithara01/ecommerce/utils/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	SearchProducts(query string) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) SearchProducts(query string) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("name LIKE ?", "%"+query+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// 	"gorm.io/gorm"

// 	"github.com/fathimasithara01/ecommerce/database"
// 	"github.com/fathimasithara01/ecommerce/utils/models"
// )

// func CreateProduct(product models.Product) (*models.Product, error) {
// 	err := database.DB.Create(&product).Error
// 	return &product, err
// }

// func GetAllProduct() ([]models.Product, error) {
// 	var products []models.Product
// 	err := database.DB.Find(&products).Error
// 	return products, err
// }

// func UpdateProduct(id uint, update models.ProductUpdate) (*models.Product, error) {
// 	var product models.Product

// 	if err := database.DB.First(&product, id).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("product not found")
// 		}
// 		return nil, err
// 	}

// 	product.Name = update.Name
// 	product.Description = update.Description
// 	product.Price = update.Price
// 	product.Stock = update.Stock

// 	err := database.DB.Save(&product).Error
// 	return &product, err
// }

// func DeleteProduct(id uint) error {
// 	var product models.Product
// 	return database.DB.Delete(&product, id).Error
// }

// func SearchProduct(query string) ([]models.Product, error) {
// 	var products []models.Product
// 	err := database.DB.Where("name LIKE ?", "%"+query+"%").Find(&products).Error
// 	return products, err
// }
