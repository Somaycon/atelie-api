package handler

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show product
// @Description Show a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} GetProductByIdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product/{id} [get]
func GetProductByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	product := schemas.Products{}

	if err := db.First(&product, id).Error; err != nil{
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return 
	}
	sendSucess(ctx, "get-product-by-id", product)
}