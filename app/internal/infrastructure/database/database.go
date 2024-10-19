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

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo", user, pass, host, port, dbname)
	var err error
	max := 3
	for i := 0; i < max; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
		if err == nil {
			return db, nil
		}
		con, _ := db.DB()
		if err := con.Ping(); err != nil {
			sugar.Infof("DB接続に%d回失敗しました。5秒後に再実行します。:%s", i, err)
			time.Sleep(5 * time.Second)
		}
	}

	return nil, fmt.Errorf("DB接続に%d回以上失敗しました。接続情報を確認してください: %w", max, err)
}
