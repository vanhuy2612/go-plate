package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/service"
)

type ISystemController interface {
	HandleCPUBound()
	HandleIOReadDB()
	HandleIOPublish2Queue()
	HandleIOCallExternalApi()
}

type SystemController struct {
	Service     *service.SystemService
	UserService *service.UserService
	CpuService  *service.CpuService
}

func (sc *SystemController) HandleCpuBound(c *gin.Context) {
	sc.CpuService.ResizeImage(c)
}

func (sc *SystemController) HandleIOReadDB(c *gin.Context) {
	sc.UserService.Find(c)
}

func (sc *SystemController) HandleIOWriteDB(c *gin.Context) {
	sc.UserService.Save(c)
}

func (sc *SystemController) HandleIOPublish2Queue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (sc *SystemController) HandleIOCallExternalApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}
