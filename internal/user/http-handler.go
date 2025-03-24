package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin-app/internal/server"
)

type HttpHandler struct {
	service Servicer
}

func NewHttpHandler(service Servicer) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) CreateUser(c *gin.Context) {
	user, err := NewOuterUser(c)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	id, err := h.service.CreateUser(c.Request.Context(), user)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, id)
}

func (h *HttpHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers(c.Request.Context())
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, users)
}

func (h *HttpHandler) GetUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	user, err := h.service.GetUserById(c.Request.Context(), id)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, user)
}

func (h *HttpHandler) UpdateUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	user, err := NewOuterUser(c)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	if err = h.service.UpdateUserById(c.Request.Context(), id, user); err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c)
}

func (h *HttpHandler) DeleteUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	if err = h.service.DeleteUserById(c.Request.Context(), id); err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c)
}

func (h *HttpHandler) IsUserExistsByUsername(c *gin.Context) {
	username := c.Param("username")

	isExists, err := h.service.IsUserExistsByUsername(c.Request.Context(), username)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, isExists)
}
