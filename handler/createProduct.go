package handler

import (
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Request body"
// @Success 200 {object} CreateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [post]
func CreateProductHandler(ctx *gin.Context) {
	request := CreateProductRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product := schemas.Products{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
	}

	if err := db.Create(&product).Error; err != nil {
		logger.Errorf("Error creating product: %v", err)
		sendError(ctx, http.StatusInternalServerError,"error creating product on database")
		ctx.JSON(500, gin.H{"error": "Error creating product"})
		return
	}
	sendSucess(ctx, "create-product", product)
}