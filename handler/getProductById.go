package handler

import "github.com/gin-gonic/gin"

func GetProductByIdHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mesage": "GET do produto com id",
	})
}