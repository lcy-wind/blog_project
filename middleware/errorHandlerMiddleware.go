package middleware

import (
	loggerutils "blog_project/loggerUtils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppError struct {
	StatusCode int    // HTTP状态码
	Code       string // 错误代码
	Message    string // 错误消息
	Err        error  // 原始错误
}

// 实现error接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError 创建新的应用错误
func NewAppError(statusCode int, code, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

var logger = loggerutils.Logger

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 先执行后续的处理器（路由函数、其他中间件）
		c.Next()

		// 2. 检查是否有错误（只有在有错误时才处理）
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			fmt.Printf("捕获到错误信息: %v\n", err.Err)

			// 判断错误类型
			switch e := err.Err.(type) {
			case *AppError:
				// 日志已初始化，可安全使用
				logger.Error("应用错误",
					zap.String("路径", c.Request.URL.Path),
					zap.Int("状态码", e.StatusCode),
					zap.String("消息", e.Message),
					zap.Error(e.Err))

				// 返回错误响应
				c.JSON(e.StatusCode, gin.H{
					"error": e.Message,
					"code":  e.Code,
				})

			default:
				logger.Error("未知错误",
					zap.String("路径", c.Request.URL.Path),
					zap.Error(err.Err))

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "服务器内部错误",
					"code":  "INTERNAL_ERROR",
				})
			}

			// 只有有错误时才终止后续处理
			c.Abort()
		}
	}
}
