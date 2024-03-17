package pkg

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// loggernの初期化
func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// ロガーをコンテキストから取得
func GetLogger(c *gin.Context) (*zap.Logger, error) {
	l, exists := c.Get("logger")
	if !exists {
		return nil, errors.New("ロガーが存在しません")
	}
	logger, ok := l.(*zap.Logger)
	if !ok {
		return nil, errors.New("ロガーが不正です")
	}
	return logger, nil
}
