package main

import (
	"github.com/TTT0420/golangCleanArch/internal/infrastructure"
	"github.com/TTT0420/golangCleanArch/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logger, err := pkg.NewLogger()
	if err != nil {
		return
	}
	defer logger.Sync()
	r.Use(pkg.LoggingReq(logger))
	infrastructure.SetupRoutes(r, logger)
	r.Run()
}
