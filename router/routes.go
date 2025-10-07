package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/product/:id", func(ctx *gin.Context){
			id := ctx.Param("id")
			ctx.JSON(200, gin.H{
				"mesage" : "GET do produto com id: " + id,
			})
		})
		v1.GET("/products", func (ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "GET de produtos!",
			})
		})
		v1.POST("/product", func(ctx *gin.Context){
			ctx.JSON(201, gin.H{
				"message": "POST de produto!",
			})
		})
		v1.PUT("/product/:id", func(ctx *gin.Context){
			id := ctx.Param("id")
			ctx.JSON(200, gin.H{
				"message": "PUT do produto com id: " + id,
			})
		})
		v1.DELETE("/product/:id", func(ctx *gin.Context){
			id := ctx.Param("id")
			ctx.JSON(200, gin.H{
				"message": "DELETE do produto com id: " + id,
			})
		})
	}
}