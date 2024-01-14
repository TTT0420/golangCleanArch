package handler

import (
	"net/http"

	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostUsecase usecase.PostUsecase
}

func NewPostController(postUsecase usecase.PostUsecase) *PostController {
	return &PostController{
		PostUsecase: postUsecase,
	}
}

func (h *PostController) GetAllPosts(c *gin.Context) {
	posts, err := h.PostUsecase.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": posts})
}
