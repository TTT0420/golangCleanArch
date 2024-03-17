package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDB(logging *zap.Logger) (*gorm.DB, error) {
	sugar := logging.Sugar()
	stdLogger := log.New(os.Stdout, "\r\n", log.LstdFlags)
	newLogger := logger.New(
		stdLogger,
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := "root:root@tcp(mysql:3306)/mate?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	for i := 0; i < 3; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
		if err == nil {
			return db, nil
		}
		sugar.Infof("DB接続に%d回失敗しました。5秒後に再実行します。:%s", i, err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("DB接続に3回以上失敗しました。接続情報を確認してください: %s", err)
}
