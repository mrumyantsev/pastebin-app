package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin-app/internal/server"
	"github.com/mrumyantsev/pastebin-app/internal/user"
)

type HttpHandler struct {
	service Servicer
}

func NewHttpHandler(service Servicer) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) SignUp(c *gin.Context) {
	user, err := user.NewOuterUser(c)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	token, err := h.service.SignUp(c.Request.Context(), user)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, token)
}

func (h *HttpHandler) SignIn(c *gin.Context) {
	auth, err := NewOuterAuth(c)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	token, err := h.service.SignIn(c.Request.Context(), auth)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, token)
}
