package main

import (
	"github.com/TTT0420/golangCleanArch/internal/interface/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)
	r.Run(":8080")
}
