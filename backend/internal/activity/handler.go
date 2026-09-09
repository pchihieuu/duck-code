// handler.go
package activity

import (
	"net/http"
	"strconv"

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
	rg.GET("/activity/me", h.listMine)
}

func (h *Handler) listMine(c *gin.Context) {
	userID := middleware.GetUserID(c)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	items, err := h.svc.ListRecent(c.Request.Context(), userID, limit)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, items)
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}