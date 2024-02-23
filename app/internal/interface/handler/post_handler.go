package handler

import (
	"fmt"
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
		c.JSON(http.StatusBadGateway, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "created post successfully"})

}

// 投稿編集
func (h *PostHandler) EditPost(c *gin.Context) {
	if err := h.PostUsecase.EditPost(c); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "post was edited successfully"})

}

// 投稿削除
func (h *PostHandler) DeletePost(c *gin.Context) {
	if err := h.PostUsecase.DeletePost(c); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "post was deleted successfully"})

}
