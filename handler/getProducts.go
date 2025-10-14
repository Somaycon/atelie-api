package handler

import (
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List products
// @Description List all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} GetAllProductsResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func GetAllProductsHendler(ctx *gin.Context) {
	products := []schemas.Products{}

	if err := db.Find(&products).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, "error listing products")
		return 
	}
	sendSucess(ctx, "get-all-products", products)
}