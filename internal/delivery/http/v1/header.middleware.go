package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/pkg/basic"
	"go.uber.org/zap"
)

type HeaderMiddleware struct {
	basic.BaseController[any]
	logger *zap.Logger
}

func NewHeaderMiddleware(logger *zap.Logger) *HeaderMiddleware {
	return &HeaderMiddleware{
		logger: logger,
	}
}

func (m *HeaderMiddleware) validateHeaders(c *gin.Context, requiredHeaders []string) bool {
	for _, header := range requiredHeaders {
		if c.GetHeader(header) == "" {
			m.Error(c, http.StatusBadRequest, "Header not found: "+header)
			c.Abort()
			return false
		}
	}
	return true
}

func (m *HeaderMiddleware) HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !m.validateHeaders(c, []string{"X-Timestamp", "X-Signature"}) {
			return
		}

		timestamp := c.GetHeader("X-Timestamp")
		_ = c.GetHeader("X-Signature")

		systemTimestamp := time.Now().Unix()
		parsedTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			m.logger.Error("Error parsing timestamp", zap.Error(err))
			m.Error(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		m.logger.Info("X-Timestamp", zap.Int64("timestamp", parsedTimestamp))

		if parsedTimestamp < systemTimestamp-150 {
			m.Error(c, http.StatusBadRequest, "X-Timestamp is expired")
			c.Abort()
			return
		}

		c.Next()
	}
}
