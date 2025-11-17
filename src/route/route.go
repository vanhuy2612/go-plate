package route

import (
	"github.com/gin-gonic/gin"
)
import "root/src/controller"

func InitRoute(router *gin.Engine) {
	var systemCtl = &controller.SystemController{}
	api := router.Group("/api/v1")
	{
		api.GET("/cpu-bound", systemCtl.HandleCpuBound)
		api.GET("/io-read-db", systemCtl.HandleIOReadDB)
		api.GET("/io-write-db", systemCtl.HandleIOWriteDB)
		api.GET("/io-publish-2-queue", systemCtl.HandleIOPublish2Queue)
		api.GET("/io-call-external-api", systemCtl.HandleIOCallExternalApi)
	}
}
