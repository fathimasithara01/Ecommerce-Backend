package repository

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/utils/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepo{db: db}
}
func (r *categoryRepo) Create(category *models.Category) error {
	if category.CreatedBy != 0 {
		var brand models.Brand
		if err := r.db.First(&brand, category.CreatedBy).Error; err != nil {
			return errors.New("invalid brand_id")
		}
	}

	var existing models.Category
	if err := r.db.Where("LOWER(name) = ?", strings.ToLower(category.Name)).First(&existing).Error; err == nil {
		return fmt.Errorf("category %q already exists", category.Name)
	}

	if err := r.db.Create(category).Error; err != nil {
		return err
	}

	return nil
}

func (r *categoryRepo) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepo) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Brands").Find(&categories).Error
	return categories, err
}

func (r *categoryRepo) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Brands").First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &category, err
}
