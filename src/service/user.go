package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"root/src/models"
)

type IUserService interface {
	GetAll()
}

type UserService struct {
	DB *gorm.DB
}

func (u *UserService) GetAll(c *gin.Context) {
	var users []models.User
	u.DB.Where("").Limit(10).Offset(0).Find(&users)
	c.JSON(http.StatusOK, gin.H{"d": users})
}
