package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ISystemController interface {
	HandleCPUBound()
	HandleIOReadDB()
	HandleIOPublish2Queue()
	HandleIOCallExternalApi()
}

type SystemController struct {
}

func (sc *SystemController) HandleCpuBound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (sc *SystemController) HandleIOReadDB(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (sc *SystemController) HandleIOWriteDB(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (sc *SystemController) HandleIOPublish2Queue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (sc *SystemController) HandleIOCallExternalApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}
