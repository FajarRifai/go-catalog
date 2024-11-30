package services

import (
	"go-catalog/models"
	"go-catalog/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(product models.Product) (int64, error) {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repo.GetProducts()
}

func (s *ProductService) GetProductById(id int) (models.Product, error) {
	return s.Repo.GetProductById(id)
}

func (s *ProductService) UpdateProduct(product models.Product) error {
	return s.Repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repo.DeleteProduct(id)
}

func (s *ProductService) GetProductByCodes(codes []string) ([]models.Product, error) {
	return s.Repo.GetProductByCodes(codes)
}
