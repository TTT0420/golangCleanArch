package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Message interface{} `json:"message"`
}

type ErrorResponseData struct {
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}

func RespondJSON(c *gin.Context, statusCode int, message interface{}, err interface{}) {
	if statusCode >= http.StatusBadRequest {
		// エラーレスポンス
		c.JSON(statusCode, ErrorResponseData{Message: message, Error: err})
	} else {
		// 正常レスポンス
		c.JSON(statusCode, ResponseData{Message: message})
	}
}
