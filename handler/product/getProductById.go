package product

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

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