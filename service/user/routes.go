package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/login", h.handleLogin)
	r.POST("/register", h.handleRegister)
}

func (h *Handler) handleLogin(c *gin.Context) {

}

func (h *Handler) handleRegister(c *gin.Context) {

}
