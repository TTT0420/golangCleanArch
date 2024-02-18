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
	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("Ua", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Since(oldTime)),
	)
}
