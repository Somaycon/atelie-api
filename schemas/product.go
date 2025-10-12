package schemas

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name string
	Price float64
	Description string
}