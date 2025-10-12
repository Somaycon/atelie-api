package handler

import "github.com/gin-gonic/gin"

func UpdateProductHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PUT do produto com id",
	})
}