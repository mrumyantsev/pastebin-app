package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin-app/internal/reqerrors"
)

func HttpSuccessResponse(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		c.Status(http.StatusOK)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data[0],
	})
}

func HttpErrorResponse(c *gin.Context, err error) {
	SetError(c, err)

	reqErr, ok := err.(reqerrors.Error)
	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	c.JSON(reqErr.StatusCode(), gin.H{
		"error": reqErr.Error(),
	})
}
