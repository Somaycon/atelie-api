package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context){
		c.String(200, "Olá Mundo!")
	})
	server.Run()
}