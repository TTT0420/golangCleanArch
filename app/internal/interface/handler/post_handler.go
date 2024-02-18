package handler

import (
	"net/http"

	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostUsecase usecase.PostUsecase
}

func NewPostHandler(postUsecase usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		PostUsecase: postUsecase,
	}
}

// 全件取得
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.PostUsecase.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": posts})
}

// 投稿登録
func (h *PostHandler) AddPost(c *gin.Context) {

	if err := h.PostUsecase.AddPost(c); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "someting went wrong, try again"})
	}
}
