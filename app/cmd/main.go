package main

import (
	"github.com/TTT0420/golangCleanArch/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	infrastructure.SetupRoutes(r)
	r.Run(":8080")
}
