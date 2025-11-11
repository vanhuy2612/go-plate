package storage

import (
	"fmt"
	"log"
	"root/src/util"
	"time"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var DB *gorm.DB

func ConnectDatabase() {
	dsn := util.GetEnv("DB_URL")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic(">>>>>>> Failed to connect to database! <<<<<<<<")
		fmt.Printf(">>>>>>> Failed to connect to database! <<<<<<<< %s", dsn)
	} else {
		fmt.Printf("******* Database connected *********")
	}

	// ✅ Get underlying *sql.DB to configure connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("❌ Failed to get sql.DB:", err)
	}

	// ✅ Pool settings
	sqlDB.SetMaxOpenConns(100)               // max total connections
	sqlDB.SetMaxIdleConns(50)                // max idle connections
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // lifetime before recycle
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)  // idle timeout (Go 1.15+)
	if err != nil {
		return
	}

	DB = database
}
