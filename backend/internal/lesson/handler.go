// handler.go
package lesson

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

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/courses/:courseId/lessons", h.listByCourse)
	rg.GET("/lessons/:id", h.getByID)
}

func RegisterAdminRoutes(rg *gin.RouterGroup, h *Handler) {
	g := rg.Group("/admin/lessons")
	g.POST("", h.create)
	g.PATCH("/:id", h.update)
	g.DELETE("/:id", h.delete)
}

func (h *Handler) listByCourse(c *gin.Context) {
	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid course id")
		return
	}
	lessons, err := h.svc.ListByCourseID(c.Request.Context(), courseID)
	if err != nil {
		writeErr(c, err)
		return
	}
	out := make([]ListItemResponse, len(lessons))
	for i, l := range lessons {
		out[i] = ToListItem(&l)
	}
	response.OK(c, http.StatusOK, out)
}

func (h *Handler) getByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid lesson id")
		return
	}
	l, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToDetail(l))
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}
	l, err := h.svc.Create(c.Request.Context(), req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusCreated, ToDetail(l))
}

func (h *Handler) update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid lesson id")
		return
	}
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}
	l, err := h.svc.Update(c.Request.Context(), id, req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToDetail(l))
}

func (h *Handler) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid lesson id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"deleted": true})
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}