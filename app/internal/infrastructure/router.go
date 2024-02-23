package infrastructure

import (
	"github.com/TTT0420/golangCleanArch/internal/infrastructure/database"
	"github.com/TTT0420/golangCleanArch/internal/infrastructure/repository"
	"github.com/TTT0420/golangCleanArch/internal/interface/handler"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	db := database.InitializeDB()
	postRepo := repository.NewPostRepositoryImpl(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(*postUsecase)

	r.GET("/get_posts", postHandler.GetAllPosts)
	r.POST("/add_post", postHandler.AddPost)
	r.PATCH("/edit_post", postHandler.EditPost)
	r.DELETE("/delete_post", postHandler.DeletePost)

}
