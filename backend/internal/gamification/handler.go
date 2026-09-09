package gamification

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/middleware"
	apperrors "backend/pkg/errors"
	"backend/pkg/response"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/gamification/me", h.getMe)
}

func (h *Handler) getMe(c *gin.Context) {
	userID := middleware.GetUserID(c)

	status, err := h.svc.GetStatus(c.Request.Context(), userID)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, status)
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}