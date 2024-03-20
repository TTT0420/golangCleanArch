package pkg

import (
	"testing"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	dsn := "root:root@tcp(mysql:3306)/mate?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// テスト用のテーブルを作成
	if err := db.AutoMigrate(&entity.Post{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}
