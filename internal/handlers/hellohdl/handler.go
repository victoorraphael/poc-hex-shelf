package hellohdl

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/poc-hex-shelf/internal/core/ports"
)

type HTTPHandler struct {
	service ports.HelloService
}

func New(srv ports.HelloService) *HTTPHandler {
	return &HTTPHandler{service: srv}
}

func (h *HTTPHandler) Get(c *gin.Context) {
	msg := h.service.Get()
	c.JSON(200, msg)
}
