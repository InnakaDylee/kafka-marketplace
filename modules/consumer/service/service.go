package service

import (
	"errors"
	"gorm.io/gorm"
	"kafka-marketplace/modules/consumer/model"
	"kafka-marketplace/modules/consumer/repository"
)

type ConsumerService interface {
	GetAll() ([]model.Consumer, error)
	GetByID(id int) (model.Consumer, error)
	Create(consumer model.Consumer) (model.Consumer, error)
	Update(consumer model.Consumer) (model.Consumer, error)
	Delete(id int) error
}

type consumerService struct {
	repo repository.ConsumerRepository
}

func NewConsumerService(repo repository.ConsumerRepository) ConsumerService {
	return &consumerService{repo}
}

func (s *consumerService) GetAll() ([]model.Consumer, error) {
	return s.repo.GetAll()
}
func (s *consumerService) GetByID(id int) (model.Consumer, error) {
	consumer, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Consumer{}, errors.New("consumer not found")
		}
		return model.Consumer{}, err
	}
	return consumer, nil
}
func (s *consumerService) Create(consumer model.Consumer) (model.Consumer, error) {
	return s.repo.Create(consumer)
}
func (s *consumerService) Update(consumer model.Consumer) (model.Consumer, error) {
	_, err := s.repo.GetByID(consumer.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Consumer{}, errors.New("consumer not found")
		}
		return model.Consumer{}, err
	}
	return s.repo.Update(consumer)
}
func (s *consumerService) Delete(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("consumer not found")
		}
		return err
	}
	return s.repo.Delete(id)
}