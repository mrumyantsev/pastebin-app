package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin-app/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RestyLogger struct {
}

func NewRestyLogger() *RestyLogger {
	return &RestyLogger{}
}

func (l *RestyLogger) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)

		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		now := time.Now()

		c.Next()

		durationMicroseconds := time.Since(now).Microseconds()

		var logEvent *zerolog.Event

		err := server.GetError(c)
		if err != nil {
			logEvent = log.Error()
		} else {
			logEvent = log.Info()
		}

		logEvent.
			Int64("userId", server.GetUserId(c)).
			Str("ip", c.ClientIP()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Int("size", c.Writer.Size()).
			Int64("duration", durationMicroseconds).
			Bytes("body", body).
			Err(err).
			Send()
	}
}
