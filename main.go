package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		user := os.Getenv("SAMPLE_APP_USER")
		c.JSON(200, gin.H{"user": user})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "GO PATS")
	})

	router.Run(":8080")
}
