package handler

import "github.com/gin-gonic/gin"

func GetAllProductsHendler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET de produtos!",
	})
}