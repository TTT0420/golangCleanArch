package infrastructure

import (
	"log"

	"github.com/TTT0420/golangCleanArch/internal/infrastructure/database"
	"github.com/TTT0420/golangCleanArch/internal/infrastructure/repository"
	"github.com/TTT0420/golangCleanArch/internal/interfaces/handler"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/TTT0420/golangCleanArch/pkg"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupRoutes(r *gin.Engine) {
	// DB接続
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("DB接続失敗:%v", err)
		return
	}

	// バリデーションの初期化
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ContentsCheck", pkg.ContentsCheck)
	}

	// ユーザーに関する処理
	userRepo := repository.NewUserRepositoryImpl(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(*userUsecase)

	// 投稿に関する処理
	postRepo := repository.NewPostRepositoryImpl(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(*postUsecase)

	// ユーザー関連のエンドポイント

	r.GET("/get_user/:id", userHandler.GetUserByID)
	r.POST("/add_user", userHandler.AddUser)
	r.PATCH("/edit_user/:id", userHandler.EditUser)
	r.DELETE("/delete_user/:id", userHandler.DeleteUser)

	r.GET("/get_posts", postHandler.GetAllPosts)
	r.POST("/add_post", postHandler.AddPost)
	r.PATCH("/edit_post/:id", postHandler.EditPost)
	r.DELETE("/delete_post/:id", postHandler.DeletePost)
}
