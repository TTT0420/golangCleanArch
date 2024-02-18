package main

import (
	"github.com/TTT0420/golangCleanArch/internal/infrastructure"
	"github.com/TTT0420/golangCleanArch/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(pkg.LoggingRqs)
	infrastructure.SetupRoutes(r)
	r.Run(":8080")
}
