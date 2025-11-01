package core

import (
	"root/src/route"
	"root/src/storage"
	"root/src/util"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	storage.ConnectDatabase()
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	route.InitRoute(router)
	PORT := util.GetEnv("PORT")
	router.Run(":" + PORT)
}
