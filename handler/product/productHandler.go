package product

import (
	"github.com/Somaycon/atelie-api/config"
	"gorm.io/gorm"
)
var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeProductHandler() {
	logger = config.GetLogger("productHandler")
	db = config.GetDb()
}