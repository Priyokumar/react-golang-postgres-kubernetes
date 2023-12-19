package routes

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(corsMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"message": "Hello world"})
	})

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
