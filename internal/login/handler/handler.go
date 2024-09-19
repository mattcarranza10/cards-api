package handler

import (
	"net/http"

	"cards-api/internal/login"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *login.Service
}

func NewHandler(service *login.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(c *gin.Context) {
	cmd := &LoginCommand{}
	if err := c.ShouldBindJSON(cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := h.service.Login(c, cmd.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toLoginResponse(token))
}
