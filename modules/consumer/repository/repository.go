package repository

import (
	"kafka-marketplace/modules/consumer/model"
	"gorm.io/gorm"
)


type ConsumerRepository interface {
	GetAll() ([]model.Consumer, error)
	GetByID(id int) (model.Consumer, error)
	Create(consumer model.Consumer) (model.Consumer, error)
	Update(consumer model.Consumer) (model.Consumer, error)
	Delete(id int) error
}

type consumerRepository struct{
	db *gorm.DB
}

func NewConsumerRepository(db *gorm.DB) ConsumerRepository {
	return &consumerRepository{db}
}

func (r *consumerRepository) GetAll() ([]model.Consumer, error) {
	var consumers []model.Consumer
	if err := r.db.Find(&consumers).Error; err != nil {
		return nil, err
	}
	return consumers, nil
}
func (r *consumerRepository) GetByID(id int) (model.Consumer, error) {
	var consumer model.Consumer
	if err := r.db.First(&consumer, id).Error; err != nil {
		return model.Consumer{}, err
	}
	return consumer, nil
}
func (r *consumerRepository) Create(consumer model.Consumer) (model.Consumer, error) {
	if err := r.db.Create(&consumer).Error; err != nil {
		return model.Consumer{}, err
	}
	return consumer, nil
}
func (r *consumerRepository) Update(consumer model.Consumer) (model.Consumer, error) {
	if err := r.db.Save(&consumer).Error; err != nil {
		return model.Consumer{}, err
	}
	return consumer, nil
}
func (r *consumerRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Consumer{}, id).Error; err != nil {
		return err
	}
	return nil
}