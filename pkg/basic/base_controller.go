package basic

import (
	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/pkg/types"
)

type BaseController[T any] struct{}

func (BaseController[T]) OK(c *gin.Context, message string, data T) {
	c.JSON(200, types.APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func (BaseController[T]) Error(c *gin.Context, status int, message string) {
	c.JSON(status, types.APIResponse[any]{
		Success: false,
		Message: message,
	})
}
