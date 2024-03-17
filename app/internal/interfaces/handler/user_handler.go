package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/internal/interfaces/dto"
	"github.com/TTT0420/golangCleanArch/internal/usecase"
	"github.com/TTT0420/golangCleanArch/pkg"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: usecase,
	}
}

// ユーザー登録
func (h *UserHandler) AddUser(c *gin.Context) {
	var userReq dto.AddUserReq
	if err := c.BindJSON(&userReq); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: pkg.ErrMissingParam()})
		return
	}
	fmt.Printf("Requst body: %+v\n", userReq)

	user := entity.Users{
		UserID:      userReq.UserID,
		UserName:    userReq.UserName,
		UserType:    userReq.UserType,
		IsDeleted:   false,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	id, err := h.UserUseCase.AddUser(&user)
	if err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// TODO:ログ出力する
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: errors.New(pkg.ResMsgForServerError)})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})
}

// ユーザー情報取得
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO:ログ出力する
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}

	user, err := h.UserUseCase.GetUserByID(id)
	if err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.ResForUser{Result: pkg.ResOK, User: *user})
}

// ユーザー情報更新
func (h *UserHandler) EditUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO:ログ出力する
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	var user entity.Users
	if err := c.BindJSON(&user); err != nil {
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: pkg.ErrMissingParam()})
		return
	}

	// パスパラメーターのIDをセット
	user.UserID = id
	if err := h.UserUseCase.UpdateUserByID(&user); err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// TODO:ログ出力する
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: user.ID})

}

// ユーザー削除
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO:ログ出力する
		pkg.RespondJSON(c, http.StatusBadRequest, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	if err := h.UserUseCase.DeleteUserByID(id); err != nil {
		// 型アサーション
		if appErr, ok := err.(*pkg.AppError); ok {
			// カスタムエラーの場合は、関連付けられたHTTPステータスコードでレスポンス
			pkg.RespondJSON(c, appErr.Code, pkg.GeneralResponse{Result: pkg.ResNG, Error: err})
			return
		}
		// 予期しないエラーの場合は、500エラーで返す
		pkg.RespondJSON(c, http.StatusInternalServerError, pkg.GeneralResponse{Result: pkg.ResNG, Error: fmt.Errorf(pkg.ResMsgForServerError)})
		return
	}
	pkg.RespondJSON(c, http.StatusOK, pkg.GeneralResponse{Result: pkg.ResOK, ID: id})
}
