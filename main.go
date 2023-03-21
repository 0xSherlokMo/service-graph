package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	go healthCheck()
}

func healthCheck() {
	mode := gin.ReleaseMode
	if os.Getenv("env") != "prod" {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
