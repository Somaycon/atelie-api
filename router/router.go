package router

import "github.com/gin-gonic/gin"

func Initalize() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Mundo!",
		})
	})
	router.Run("8080")
}