package main

import (
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
	defer logger.Sync()
	r.Use(pkg.LoggingReq(logger))
	infrastructure.SetupRoutes(r, logger)
	r.Run()
}
