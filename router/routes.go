package router

import (
	"github.com/Somaycon/atelie-api/handler"
	"github.com/Somaycon/atelie-api/handler/product"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	
	v1 := router.Group("/api/v1")
	{
		v1.GET("/product/:id", product.GetProductByIdHandler)
		v1.GET("/products", product.GetAllProductsHendler)
		v1.POST("/product", product.CreateProductHandler)
		v1.PUT("/product/:id", product.UpdateProductHandler)
		v1.DELETE("/product/:id",product.DeleteProductHandler)
	}
}