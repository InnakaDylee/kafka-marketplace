package repository

import (
	"kafka-marketplace/modules/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	GetByID(id int) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (r *productRepository) GetByID(id int) (model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}
func (r *productRepository) Create(product model.Product) (model.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}
func (r *productRepository) Update(product model.Product) (model.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}
func (r *productRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}