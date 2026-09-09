package progress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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
	rg.POST("/lessons/:id/complete", h.completeLesson)
	rg.GET("/courses/:courseId/progress", h.getCourseProgress)
}

func (h *Handler) completeLesson(c *gin.Context) {
	userID := middleware.GetUserID(c)
	lessonID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid lesson id")
		return
	}
	if err := h.svc.CompleteLesson(c.Request.Context(), userID, lessonID); err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"completed": true})
}

func (h *Handler) getCourseProgress(c *gin.Context) {
	userID := middleware.GetUserID(c)
	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid course id")
		return
	}
	views, err := h.svc.GetCourseProgress(c.Request.Context(), userID, courseID)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, views)
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}