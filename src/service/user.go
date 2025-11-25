package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (u *UserService) Find(c *gin.Context) {
	var users []models.User
	result := u.DB.Where("name LIKE ?", "%huy%").Limit(10).Offset(0).Find(&users)
	if result.Error != nil {
		fmt.Println("Error when get data from DB")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"d": users})
}

func (u *UserService) Save(c *gin.Context) {
	id := uuid.New().String()
	user := models.User{
		Name:     "huy-" + id,
		Email:    "vanhuy2612@gmail.com",
		Password: id,
	}
	result := u.DB.Save(&user)
	if result.Error != nil {
		fmt.Println("Error when insert to DB")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
