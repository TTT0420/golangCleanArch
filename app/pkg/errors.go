package pkg

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    // HTTPステータスコード
	Message string // エラーメッセージ
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) error {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// 共通エラーメソッド。レコードがない場合
func ErrRecordNotFound(id int) error {
	fmt.Println("レコードないid: ", id)
	return NewAppError(http.StatusNotFound, fmt.Sprintf("record not found. id: %d", id))
}

// 共通エラーメソッド。パラメーター
func ErrMissingParam() error {
	return NewAppError(http.StatusBadRequest, ResMsgForInvalidReq)
}
