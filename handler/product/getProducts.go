package product

import (
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetAllProductsHendler(ctx *gin.Context) {
	products := []schemas.Products{}

	if err := db.Find(&products).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, "error listing products")
		return 
	}
	sendSucess(ctx, "get-all-products", products)
}