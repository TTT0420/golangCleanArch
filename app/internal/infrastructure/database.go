package infrastructure

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := "root:root@tcp(mysql:3306)/mate?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
