package services

import (
	repository "github.com/fathimasithara01/ecommerce/src/repository/admin"
	"github.com/fathimasithara01/ecommerce/utils/models"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repo.GetByID(id)
}
