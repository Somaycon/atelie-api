package handler

import (
	"fmt"
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

// @BasePath /api/v1

// @Summary Update product
// @Description Update a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id query string true "Product Identification"
// @Param opening body UpdateProductRequest true "Product data to Update"
// @Success 200 {object} UpdateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product/{id} [put]
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

// @BasePath /api/v1

// @Summary Delete Product
// @Description Delete a Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} DeleteProductResponse
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