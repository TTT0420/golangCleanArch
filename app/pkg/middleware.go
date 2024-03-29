package pkg

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingReq(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("リクエストを受け付けました\n",
			zap.String("User-Agent", c.GetHeader("User-Agent")),
			zap.String("path", c.Request.URL.Path))
		oldTime := time.Now()
		c.Next()
		logger.Info("返却内容\n", zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Since(oldTime)))
		logger.Info("リクエストを処理しました")
	}
}
