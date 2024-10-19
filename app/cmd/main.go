package main

import (
	"fmt"
	"log"

	"github.com/TTT0420/golangCleanArch/app/internal/infrastructure"
	"github.com/TTT0420/golangCleanArch/app/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logger, err := pkg.NewLogger()
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error(fmt.Sprintf("予期せぬエラーが発生しました.%v", err))
		}
	}()
	r.Use(pkg.LoggingReq(logger))
	infrastructure.SetupRoutes(r, logger)
	if err := r.Run(); err != nil {
		logger.Error(fmt.Sprintf("予期せぬエラーが発生しました.%v", err))
	}
}
