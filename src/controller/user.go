package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"root/src/models"
)

type IUserController interface {
	GetAll()
}

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) GetAll(c *gin.Context) {
	response := []string{"huy", "dinh", "van"}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (uc *UserController) Detail(c *gin.Context) {
	var users []models.User
	id := c.Param("id")
	fmt.Printf("id %s\n", id)
	uc.DB.Where("email LIKE ?", "huydv%").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
