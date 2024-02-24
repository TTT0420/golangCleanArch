package pkg

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingRqs(c *gin.Context) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err.Error())
	}
	oldTime := time.Now()
	ua := c.GetHeader("User-Agent")
	c.Next()
	logger.Info("リクエストを処理しました",
		zap.String("path", c.Request.URL.Path),
		zap.String("User-Agent", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Since(oldTime)),
	)
}
