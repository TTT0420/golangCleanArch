package pkg

import (
	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type GeneralResponse struct {
	Result  string        `json:"result"`
	Message string        `json:"message,omitempty"`
	ID      int           `json:"id,omitempty"`
	Posts   []entity.Post `json:"posts,omitempty"`
	Error   error         `json:"error,omitempty"`
}

func RespondJSON(c *gin.Context, statusCode int, resp GeneralResponse) {
	c.JSON(statusCode, resp)
}
