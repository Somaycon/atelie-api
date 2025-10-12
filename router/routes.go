package router

import (
	"github.com/Somaycon/atelie-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/product/:id", handler.GetProductByIdHandler)
		v1.GET("/products", handler.GetAllProductsHendler)
		v1.POST("/product", handler.CreateProductHandler)
		v1.PUT("/product/:id", handler.UpdateProductHandler)
		v1.DELETE("/product/:id",handler.DeleteProductHandler)
	}
}