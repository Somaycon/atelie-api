package product

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete Product
// @Description Delete a Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} product.DeleteProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product/{id} [delete]
func DeleteProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id ==""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	product := schemas.Products{}

		if err := db.First(&product, id).Error; err != nil{
			sendError(ctx, http.StatusNotFound, fmt.Sprintf("Product with id: %s not found", id))
			return 
		}

		if err := db.Delete(&product).Error; err != nil{
			sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting product with id: %s", id))
			return 
		}

		sendSucess(ctx, "delete-product", product)
	
}