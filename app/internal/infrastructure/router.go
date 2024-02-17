package infrastructure

import (
	"github.com/TTT0420/golangCleanArch/internal/domain/repository"
	"github.com/TTT0420/golangCleanArch/internal/infrastructure"
	"github.com/TTT0420/golangCleanArch/internal/interface/handler"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	db := infrastructure.InitializeDB()
	postRepo := repository.NewPostRepositoryImpl(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(*postUsecase)

	r.GET("/posts", postHandler.GetAllPosts)

}
