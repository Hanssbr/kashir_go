package services

import (
	"kashir_go/models"
	"kashir_go/repositories"
)

type CategoriesService struct {
	repo *repositories.CategoriesRepository
}

func NewCategoryService(repo *repositories.CategoriesRepository) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) GetAll() ([]models.Categories, error) {
	return s.repo.GetAll()
}

func (s *CategoriesService) Create(data *models.Categories) error {
	return s.repo.Create(data)
}

func (s *CategoriesService) GetByID(id int) (*models.Categories, error) {
	return s.repo.GetByID(id)
}

func (s *CategoriesService) Update(category *models.Categories) error {
	return s.repo.Update(category)
}

func (s *CategoriesService) Delete(id int) error {
	return s.repo.Delete(id)
}
