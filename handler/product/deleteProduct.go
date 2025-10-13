package product

import "github.com/gin-gonic/gin"

func DeleteProductHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE do produto com id",
	})
}