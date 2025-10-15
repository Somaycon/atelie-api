package schemas

import "gorm.io/gorm"

type MaterialsList struct {
	gorm.Model
	Product   Products `gorm:"foreignKey:ProductId"`
	Materials Materials `gorm:"foreignKey:MaterialId"`
	Quantity  int `json:"quantity"`
}

type MaterialsListResponse struct {
	ID        string `json:"id"`
	ProductId string `json:"productId"`
	MaterialId string `json:"materialId"`
	Quantity  int    `json:"quantity"`
}