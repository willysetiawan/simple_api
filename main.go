package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/check_version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0"})
	})

	r.Run()
}
