package service

import (
	"context"
	"errors"
	"kafka-marketplace/modules/payment/model"
	"kafka-marketplace/modules/payment/repository"
	"kafka-marketplace/packages/queue"
	"log"

	"gorm.io/gorm"
)
type PaymentService interface {
	GetAll() ([]model.Payment, error)
	GetByID(id int) (model.Payment, error)
	Create(payment model.Payment) (model.Payment, error)
	Update(payment model.Payment) (model.Payment, error)
	Delete(id int) error
}

type paymentService struct {
	repo     repository.PaymentRepository
	kafkaCfg *queue.KafkaConfig
}
func NewPaymentService(repo repository.PaymentRepository, kafkaCfg *queue.KafkaConfig) PaymentService {
	return &paymentService{repo, kafkaCfg}
}

func (s *paymentService) GetAll() ([]model.Payment, error) {
	return s.repo.GetAll()
}
func (s *paymentService) GetByID(id int) (model.Payment, error) {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Payment{}, errors.New("payment not found")
		}
		return model.Payment{}, err
	}
	return payment, nil
}
func (s *paymentService) Create(payment model.Payment) (model.Payment, error) {
	return s.repo.Create(payment)
}
func (s *paymentService) Update(payment model.Payment) (model.Payment, error) {
	// cek payment
	_, err := s.repo.GetByID(payment.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Payment{}, errors.New("payment not found")
		}
		return model.Payment{}, err
	}

	updatedPayment, err := s.repo.Update(payment)
	if err != nil {
		return model.Payment{}, err
	}

	if updatedPayment.Status == "SUCCESS" {
		err = s.kafkaCfg.Write(context.Background(), updatedPayment)
		if err != nil {
			log.Printf("error writing to kafka: %v", err)
		} else {
			log.Printf("Kafka: published payment ID %d", updatedPayment.ID)
		}
	}

	return updatedPayment, nil
}
func (s *paymentService) Delete(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("payment not found")
		}
		return err
	}
	return s.repo.Delete(id)
}