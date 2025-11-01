package route

import (
	"github.com/gin-gonic/gin"
	"root/src/service"
	"root/src/storage"
)
import "root/src/controller"

var postSer = &service.PostService{DB: storage.DB}
var userCtl = &controller.UserController{DB: storage.DB}
var postCtl = &controller.PostController{Service: postSer}

func InitRoute(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test successful",
			})
		})
		api.GET("/users", userCtl.GetAll)
		api.GET("/users/:id", userCtl.Detail)

		api.GET("post", postCtl.GetAll)
	}
}
