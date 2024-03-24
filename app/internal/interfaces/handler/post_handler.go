package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TTT0420/golangCleanArch/app/internal/interfaces/dto"
	"github.com/TTT0420/golangCleanArch/app/internal/usecase"
	"github.com/TTT0420/golangCleanArch/app/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PostHandler struct {
	PostUsecase usecase.PostUsecase
	Logger      *zap.Logger
}

func NewPostHandler(postUsecase usecase.PostUsecase, logger *zap.Logger) *PostHandler {
	return &PostHandler{
		PostUsecase: postUsecase,
		Logger:      logger,
	}
}

// 全件取得
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	logger, err := pkg.GetLogger(c)
	if err != nil {
		log.Printf("error: %v", err)
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	posts, err := h.PostUsecase.GetAllPosts()
	if err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, Posts: posts})

}

// 投稿登録
func (h *PostHandler) AddPost(c *gin.Context) {
	logger, err := pkg.GetLogger(c)
	if err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	var post dto.AddPostReq
	if err := c.ShouldBindJSON(&post); err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
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
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})

}

// 投稿編集
func (h *PostHandler) EditPost(c *gin.Context) {
	logger, err := pkg.GetLogger(c)
	if err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}

	var post dto.EditPostReq
	if err := c.ShouldBindJSON(&post); err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
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
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})

}

// 投稿削除
func (h *PostHandler) DeletePost(c *gin.Context) {
	logger, err := pkg.GetLogger(c)
	if err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
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
		logger.Error(pkg.LogMsgForServerError, zap.Error(err))
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})

}
