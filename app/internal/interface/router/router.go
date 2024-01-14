package router

import (
	"github.com/TTT0420/golangCleanArch/internal/infrastructure/repository"
	"github.com/TTT0420/golangCleanArch/internal/interface/controller"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	db := repository.InitializeDB()
	postRepo := repository.NewPostRepositoryImpl(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := controller.NewPostController(*postUsecase)

	r.GET("/posts", postHandler.GetAllPosts)

}
