package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	model_log "github.com/luyanakat/golang-base-project/app/models/logmodel"
	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.SugaredLogger
}

func NewMiddleware(logger *zap.SugaredLogger) *Middleware {
	return &Middleware{logger: logger}
}

func (m *Middleware) RequestTraceMiddleware(c *gin.Context) {
	requestID := uuid.New().String()

	log := model_log.MiddlewareLogRequest{
		RequestID:   requestID,
		Method:      c.Request.Method,
		URL:         c.Request.URL.String(),
		UserAgent:   c.Request.UserAgent(),
		ContentType: c.Request.Header.Get("Content-Type"),
	}

	logJson, _ := json.Marshal(log)
	m.logger.Infof("Request: %v", string(logJson))

	c.Next()
}
