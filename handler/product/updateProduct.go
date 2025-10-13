package product

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateProductHandler(ctx *gin.Context) {
	request := UpdateProductRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil{
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	if id =="" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	
	product := schemas.Products{}
	
	if err := db.First(&product, id).Error; err != nil{
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return
	}

	if request.Name != ""{
		product.Name = request.Name
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	if request.Description != ""{
		product.Description = request.Description
	}

	if err := db.Save(&product).Error; err != nil{
		logger.Errorf("error updating product: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating product")
		return
	}
	sendSucess(ctx, "update-product", product)

}