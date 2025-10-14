package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Materials struct {
	gorm.Model
	Name         string
	Quantity     int
	PricePerUnit float64
}

type MaterialResponse struct {
	ID           string     `json:"id"`
	CreatedAt    *time.Time     `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	Name         string     `json:"name"`
	Quantity     int        `json:"quantity"`
	PricePerUnit float64    `json:"pricePerUnit"`
}