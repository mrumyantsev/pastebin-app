package middleware

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin-app/internal/jsonclean"
)

type JsonBodyCleaner struct {
}

func NewJsonBodyCleaner() *JsonBodyCleaner {
	return &JsonBodyCleaner{}
}

func (j *JsonBodyCleaner) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)

		body = jsonclean.Clean(body)

		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		c.Next()
	}
}
