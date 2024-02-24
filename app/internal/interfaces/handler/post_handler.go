package handler

import (
	"net/http"
	"strconv"

	"github.com/TTT0420/golangCleanArch/internal/interfaces/dto"
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
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, Posts: posts})

}

// 投稿登録
func (h *PostHandler) AddPost(c *gin.Context) {
	var post dto.AddPostReq
	if err := c.ShouldBindJSON(&post); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: pkg.ErrMissingParam()})
		return
	}

	id, err := h.PostUsecase.AddPost(post)
	if err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})

}

// 投稿編集
func (h *PostHandler) EditPost(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}

	var post dto.EditPostReq
	if err := c.ShouldBindJSON(&post); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	// パスパラメーターのIDをセット
	post.ID = id
	if err := h.PostUsecase.EditPostByID(post); err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})

}

// 投稿削除
func (h *PostHandler) DeletePost(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}

	if err := h.PostUsecase.DeletePostByID(id); err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})

}
