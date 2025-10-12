package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name string
	Price float64
	Description string
}

type ProductResponse struct{
	Id uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}