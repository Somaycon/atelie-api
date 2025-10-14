package router

import (
	docs "github.com/Somaycon/atelie-api/docs"
	"github.com/Somaycon/atelie-api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath)
	{
		v1.GET("/product/:id", handler.GetProductByIdHandler)
		v1.GET("/products", handler.GetAllProductsHendler)
		v1.POST("/product", handler.CreateProductHandler)
		v1.PUT("/product/:id", handler.UpdateProductHandler)
		v1.DELETE("/product/:id",handler.DeleteProductHandler)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}