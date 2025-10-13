package handler

import (
	"github.com/Somaycon/atelie-api/config"
	"github.com/Somaycon/atelie-api/handler/product"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDb()
	product.InitializeProductHandler()
}