package router

import (
	"github.com/TTT0420/golangCleanArch/internal/interface/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, postHandler handler.PostHandler) {
	r.GET("/posts", postHandler.GetAllPosts)
}
