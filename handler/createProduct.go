package handler

import "github.com/gin-gonic/gin"

func CreateProductHandler(ctx *gin.Context) {
	ctx.JSON(201, gin.H{
		"message": "POST de produto!",
	})
}