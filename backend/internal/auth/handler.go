package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/user"
	apperrors "backend/pkg/errors"
	"backend/pkg/response"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes mounts the public /auth endpoints (no auth middleware).
func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	g := rg.Group("/auth")
	g.POST("/register", h.register)
	g.POST("/login", h.login)
	g.POST("/refresh", h.refresh)
	g.POST("/logout", h.logout)
}

func (h *Handler) register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	u, err := h.svc.Register(c.Request.Context(), req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusCreated, user.ToPublicResponse(u))
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	u, tokens, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{
		"user":   user.ToPublicResponse(u),
		"tokens": tokens,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	tokens, err := h.svc.Refresh(c.Request.Context(), req.RefreshToken)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, tokens)
}

func (h *Handler) logout(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	if err := h.svc.Logout(c.Request.Context(), req.RefreshToken); err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "logged out"})
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}
