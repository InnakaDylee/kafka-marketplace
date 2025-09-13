package queue

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	consumerService "kafka-marketplace/modules/consumer/service"
	paymentModel "kafka-marketplace/modules/payment/model"
	productService "kafka-marketplace/modules/product/service"
)

type KafkaConfig struct {
	Address       []string
	Topic         string
	GroupID       string
	Batch         int
	BatchTimeout  int
	RequiredAcks  int
	Async         bool

	writer *kafka.Writer
	reader *kafka.Reader
	mu     sync.Mutex
}

// =======================
// Setup Writer
// =======================
func (k *KafkaConfig) newWriter() {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.writer == nil {
		k.writer = &kafka.Writer{
			Addr:         kafka.TCP(k.Address...),
			Topic:        k.Topic,
			Balancer:     &kafka.LeastBytes{},
			BatchSize:    k.Batch,
			BatchTimeout: time.Duration(k.BatchTimeout) * time.Millisecond,
			RequiredAcks: kafka.RequiredAcks(k.RequiredAcks),
			Async:        k.Async,
			Compression:  kafka.Lz4,
		}
	}
}

func (k *KafkaConfig) newReader() {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.reader == nil {
		k.reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers:  k.Address,
			GroupID:  k.GroupID,
			Topic:    k.Topic,
			MinBytes: 10e3,  // 10KB
			MaxBytes: 10e6,
		})
	}
}

func (k *KafkaConfig) Write(ctx context.Context, payment paymentModel.Payment) error {
	k.newWriter()

	payload, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte("payment_success"),
		Value: payload,
	}

	return k.writer.WriteMessages(ctx, msg)
}

func (k *KafkaConfig) Read(ctx context.Context, consumerSvc consumerService.ConsumerService, productSvc productService.ProductService) {
	k.newReader()

	for {
		m, err := k.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading kafka message: %v", err)
			continue
		}

		var p paymentModel.Payment
		if err := json.Unmarshal(m.Value, &p); err != nil {
			log.Printf("error unmarshal payment: %v", err)
			continue
		}

		log.Printf("Kafka: processing payment ID %d", p.ID)

		consumer, err := consumerSvc.GetByID(p.ConsumerID)
		if err == nil {
			consumer.Saldo -= p.Amount
			if _, err := consumerSvc.Update(consumer); err != nil {
				log.Println("error updating consumer:", err)
			}
		}

		product, err := productSvc.GetByID(p.ProductID)
		if err == nil {
			product.Stock -= 1
			if _, err := productSvc.Update(product); err != nil {
				log.Println("error updating product:", err)
			}
		}

		log.Printf("Kafka: payment %d processed", p.ID)
	}
}

func (k *KafkaConfig) Close() error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.writer != nil {
		if err := k.writer.Close(); err != nil {
			log.Println("error closing writer:", err)
		}
		k.writer = nil
	}

	if k.reader != nil {
		if err := k.reader.Close(); err != nil {
			log.Println("error closing reader:", err)
		}
		k.reader = nil
	}

	return nil
}
