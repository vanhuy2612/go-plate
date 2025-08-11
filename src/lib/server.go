package lib

import (
	"root/src/route"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	route.InitRoute(router)
	router.Run(":8080")
}
