package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Materials struct {
	gorm.Model
	Name         string
	CurrentStock     int
	PricePerUnit float64
}

type MaterialResponse struct {
	ID           string     `json:"id"`
	CreatedAt    *time.Time     `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	Name         string     `json:"name"`
	CurrentStock     int        `json:"currentStock"`
	PricePerUnit float64    `json:"pricePerUnit"`
}