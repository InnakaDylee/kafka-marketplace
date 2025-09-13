package model

import (
	"time"
)

type Product struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Price     int       `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}