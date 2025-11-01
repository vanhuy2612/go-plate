package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/service"
)

type IPostController interface {
	GetAll()
}

type PostController struct {
	Service *service.PostService
}

func (uc *PostController) GetAll(c *gin.Context) {
	post := service.Index()
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func (uc *PostController) Detail(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": ""})
}
