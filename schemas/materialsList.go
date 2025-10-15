package schemas

import "gorm.io/gorm"

type MaterialsList struct {
	gorm.Model
	ProductID  uint     `json:"productId"`
	Product    Products `gorm:"foreignKey:ProductID"`
	MaterialID uint     `json:"materialId"`
	Materials  Materials `gorm:"foreignKey:MaterialID"`
	Quantity  int `json:"quantity"`
}

type MaterialsListResponse struct {
	ID        string `json:"id"`
	ProductId string `json:"productId"`
	MaterialId string `json:"materialId"`
	Quantity  int    `json:"quantity"`
}