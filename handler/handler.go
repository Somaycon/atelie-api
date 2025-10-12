package handler

import (
	"github.com/gin-gonic/gin"
)

func GetAllProductsHendler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET de produtos!",
	})
}

func GetProductByIdHandler(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"mesage" : "GET do produto com id",
	})
}

func CreateProductHandler(ctx *gin.Context) {
		ctx.JSON(201, gin.H{
		"message": "POST de produto!",
	})
}

func UpdateProductHandler(ctx *gin.Context){
	ctx.JSON(200, gin.H{
				"message": "PUT do produto com id",
			})
}

func DeleteProductHandler(ctx *gin.Context){
	ctx.JSON(200, gin.H{
				"message": "DELETE do produto com id",
			})
}