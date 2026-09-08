package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	apperrors "backend/pkg/errors"
	"backend/pkg/response"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes mounts /profile/me on an already-authenticated group.
func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	g := rg.Group("/profile")
	g.GET("/me", h.getMe)
	g.PATCH("/me", h.updateMe)
}

func (h *Handler) getMe(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	p, err := h.svc.GetOrCreate(c.Request.Context(), userID)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToResponse(p))
}

func (h *Handler) updateMe(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	p, err := h.svc.Update(c.Request.Context(), userID, req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToResponse(p))
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}
