package infrastructure

import (
	"log"

	"github.com/TTT0420/golangCleanArch/internal/infrastructure/database"
	"github.com/TTT0420/golangCleanArch/internal/infrastructure/repository"
	"github.com/TTT0420/golangCleanArch/internal/interfaces/handler"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// DB接続
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("DB接続失敗:%v", err)
		return
	}
	postRepo := repository.NewPostRepositoryImpl(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(*postUsecase)

	r.GET("/get_posts", postHandler.GetAllPosts)
	r.POST("/add_post", postHandler.AddPost)
	r.PATCH("/edit_post/:id", postHandler.EditPost)
	r.DELETE("/delete_post/:id", postHandler.DeletePost)
}
