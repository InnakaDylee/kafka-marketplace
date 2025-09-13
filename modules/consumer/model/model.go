package model

import (
	"time"
)

type Consumer struct {
	ID             int         `json:"id" gorm:"primaryKey"`
	Name           string      `json:"name" gorm:"not null"`
	Saldo          int         `json:"saldo" gorm:"not null;default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}