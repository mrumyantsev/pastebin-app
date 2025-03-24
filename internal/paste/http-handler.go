package paste

import (
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

func (h *HttpHandler) CreatePaste(c *gin.Context) {
	paste, err := NewOuterPaste(c)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	id, err := h.service.CreatePaste(c.Request.Context(), paste)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, id)
}

func (h *HttpHandler) GetAllPastes(c *gin.Context) {
	pastes, err := h.service.GetAllPastes(c.Request.Context())
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, pastes)
}

func (h *HttpHandler) GetPasteById(c *gin.Context) {
	id := c.Param("base64-id")

	paste, err := h.service.GetPasteById(c.Request.Context(), id)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, paste)
}

func (h *HttpHandler) UpdatePasteById(c *gin.Context) {
	id := c.Param("base64-id")

	paste, err := NewOuterPaste(c)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	if err = h.service.UpdatePasteById(c.Request.Context(), id, paste); err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c)
}

func (h *HttpHandler) DeletePasteById(c *gin.Context) {
	id := c.Param("base64-id")

	err := h.service.DeletePasteById(c.Request.Context(), id)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c)
}

func (h *HttpHandler) IsPasteContentExistsById(c *gin.Context) {
	id := c.Param("base64-id")

	isExists, err := h.service.IsPasteContentExistsById(c, id)
	if err != nil {
		server.HttpErrorResponse(c, err)

		return
	}

	server.HttpSuccessResponse(c, isExists)
}
