package main

import (
	"github.com/TTT0420/golangCleanArch/internal/infrastructure/repository"
	"github.com/TTT0420/golangCleanArch/internal/interface/handler"
	"github.com/TTT0420/golangCleanArch/internal/interface/router"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.InitializeDB()
	postRepo := repository.NewPostRepositoryImpl(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(*postUsecase)

	r := gin.Default()
	router.SetupRoutes(r, *postHandler)
	r.Run(":8080")
}
