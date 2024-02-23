package handler

import (
	"net/http"
	"strconv"

	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
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
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{pkg.ResMsg: pkg.ResNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{pkg.ResMsg: pkg.ResOK, pkg.ResPosts: posts}, nil)

}

// 投稿登録
func (h *PostHandler) AddPost(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, gin.H{pkg.ResMsg: pkg.ResNG}, err)
		return
	}

	id, err := h.PostUsecase.AddPost(post)
	if err != nil {
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{pkg.ResMsg: pkg.ResNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{pkg.ResMsg: pkg.ResOK, pkg.ResID: id}, nil)

}

// 投稿編集
func (h *PostHandler) EditPost(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, gin.H{pkg.ResMsg: pkg.ResNG}, nil)
		return
	}

	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, gin.H{pkg.ResMsg: pkg.ResNG}, err)
		return
	}
	// パスパラメーターのIDをセット
	post.ID = id
	if err := h.PostUsecase.EditPostByID(post); err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, gin.H{pkg.ResMsg: appErr.Message}, nil)
			return
		}
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, gin.H{pkg.ResMsg: pkg.ResMsgInternalServerErr}, nil)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{pkg.ResMsg: pkg.ResOK, pkg.ResID: id}, nil)

}

// 投稿削除
func (h *PostHandler) DeletePost(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, gin.H{pkg.ResMsg: pkg.ResNG}, nil)
		return
	}

	if err := h.PostUsecase.DeletePostByID(id); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, gin.H{pkg.ResMsg: pkg.ResNG}, err)
		return
	}
	pkg.RespondJSON(c, http.StatusOK, gin.H{pkg.ResMsg: pkg.ResOK, pkg.ResID: id}, nil)

}
