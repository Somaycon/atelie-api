package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, status int, message string) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(status, gin.H{
		"message" :message,
		"status": status,
	})
}

func sendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s sucess", op),
		"data": data,
	})
}