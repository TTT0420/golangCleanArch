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

func ErrRecordNotFound(id int) error {
	return NewAppError(http.StatusNotFound, fmt.Sprintf("record not found. id: %d", id))
}

func ErrValidation(param string) error {
	return NewAppError(http.StatusBadRequest, fmt.Sprintf("param %s is invalid", param))
}
