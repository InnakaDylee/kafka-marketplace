package model

import (
	"time"
)

type Payment struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ConsumerID int    `json:"consumer_id" gorm:"not null"`
	ProductID  int    `json:"product_id" gorm:"not null"`
	Amount     int    `json:"amount" gorm:"not null"`
	Status     string `json:"status" gorm:"not null;default:'pending'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}