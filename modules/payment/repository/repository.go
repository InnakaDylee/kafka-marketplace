package repository

import (
	"kafka-marketplace/modules/payment/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetAll() ([]model.Payment, error)
	GetByID(id int) (model.Payment, error)
	Create(payment model.Payment) (model.Payment, error)
	Update(payment model.Payment) (model.Payment, error)
	Delete(id int) error
}

type paymentRepository struct {
	db *gorm.DB
}
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) GetAll() ([]model.Payment, error) {
	var payments []model.Payment
	if err := r.db.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
func (r *paymentRepository) GetByID(id int) (model.Payment, error) {
	var payment model.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}
func (r *paymentRepository) Create(payment model.Payment) (model.Payment, error) {
	if err := r.db.Create(&payment).Error; err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}
func (r *paymentRepository) Update(payment model.Payment) (model.Payment, error) {
	if err := r.db.Save(&payment).Error; err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}
func (r *paymentRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Payment{}, id).Error; err != nil {
		return err
	}
	return nil
}