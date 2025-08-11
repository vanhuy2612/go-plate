package route

import "github.com/gin-gonic/gin"

func InitRoute(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test successful",
			})
		})
	}
}
