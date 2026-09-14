package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/internal/config"
	"backend/internal/user"
	apperrors "backend/pkg/errors"
	"backend/pkg/response"
)

type Handler struct {
	svc Service
	cfg *config.Config
}

func NewHandler(svc Service, cfg *config.Config) *Handler {
	return &Handler{svc: svc, cfg: cfg}
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

	setRefreshCookie(c, h.cfg, tokens.RefreshToken, refreshMaxAgeSeconds(h.cfg))
	response.OK(c, http.StatusOK, gin.H{
		"user":   user.ToPublicResponse(u),
		"tokens": ToAccessTokenResponse(tokens), // refresh_token NEVER in the body — cookie only
	})
}

// refresh reads the refresh token from the HttpOnly cookie ONLY — there is
// no JSON body anymore (checklist item: "Không còn yêu cầu {refreshToken}
// cho /auth/refresh"). A missing or empty cookie is rejected before the
// service is even called.
func (h *Handler) refresh(c *gin.Context) {
	token, err := c.Cookie(refreshCookieName)
	if err != nil || token == "" {
		response.Err(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing refresh token cookie")
		return
	}

	tokens, err := h.svc.Refresh(c.Request.Context(), token)
	if err != nil {
		writeErr(c, err)
		return
	}

	setRefreshCookie(c, h.cfg, tokens.RefreshToken, refreshMaxAgeSeconds(h.cfg))
	response.OK(c, http.StatusOK, ToAccessTokenResponse(tokens))
}

// logout is idempotent at the HTTP layer too: a missing cookie is treated
// as "already logged out" (200), not an error — the goal state is already
// achieved. The cookie is always cleared regardless.
func (h *Handler) logout(c *gin.Context) {
	token, err := c.Cookie(refreshCookieName)
	if err == nil && token != "" {
		if err := h.svc.Logout(c.Request.Context(), token); err != nil {
			writeErr(c, err)
			return
		}
	}

	clearRefreshCookie(c, h.cfg)
	response.OK(c, http.StatusOK, gin.H{"message": "logged out"})
}

func refreshMaxAgeSeconds(cfg *config.Config) int {
	return cfg.JWTRefreshTTLHours * int(time.Hour.Seconds())
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}
