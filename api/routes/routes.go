package routes

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "This is great")
	})

	return r
}
