package handler

import (
	"net/http"

	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/TTT0420/golangCleanArch/pkg"
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
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{"message": pkg.ResponseNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{"message": pkg.ResponseOK, "posts": posts}, nil)

}

// 投稿登録
func (h *PostHandler) AddPost(c *gin.Context) {

	id, err := h.PostUsecase.AddPost(c)
	if err != nil {
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{"message": pkg.ResponseNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{"message": pkg.ResponseOK, "id": id}, nil)

}

// 投稿編集
func (h *PostHandler) EditPost(c *gin.Context) {
	id, err := h.PostUsecase.EditPost(c)
	if err != nil {
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{"message": pkg.ResponseNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{"message": pkg.ResponseOK, "id": id}, nil)

}

// 投稿削除
func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := h.PostUsecase.DeletePost(c)
	if err != nil {
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{"message": pkg.ResponseNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{"message": pkg.ResponseOK, "id": id}, nil)

}
