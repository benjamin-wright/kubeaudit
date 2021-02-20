package server

import (
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run("0.0.0.0:80")
}
