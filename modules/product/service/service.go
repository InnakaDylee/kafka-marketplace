package service

import (
	"errors"
	"gorm.io/gorm"
	"kafka-marketplace/modules/product/model"
	"kafka-marketplace/modules/product/repository"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	GetByID(id int) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(id int) error
}

type productService struct {
	repo repository.ProductRepository
}
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}
func (s *productService) GetByID(id int) (model.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}
	return product, nil
}
func (s *productService) Create(product model.Product) (model.Product, error) {
	return s.repo.Create(product)
}
func (s *productService) Update(product model.Product) (model.Product, error) {
	_, err := s.repo.GetByID(product.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}
	return s.repo.Update(product)
}
func (s *productService) Delete(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	return s.repo.Delete(id)
}