package storage

import (
	"fmt"
	"root/src/util"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var DB *gorm.DB

func ConnectDatabase() {
	dsn := util.GetEnv("DB_URL")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic(">>>>>>> Failed to connect to database! <<<<<<<<")
		fmt.Printf(">>>>>>> Failed to connect to database! <<<<<<<<")
	} else {
		fmt.Printf("******* Database connected *********")
	}

	if err != nil {
		return
	}

	DB = database
}
