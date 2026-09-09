// handler.go
package course

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

// RegisterRoutes: public, GET only.
func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/courses", h.listPublished)
	rg.GET("/courses/:courseId", h.getBySlug)
}

// RegisterAdminRoutes: caller (router.go) wraps rg with
// RequireAuth + RequireRole("admin").
func RegisterAdminRoutes(rg *gin.RouterGroup, h *Handler) {
	g := rg.Group("/admin/courses")
	g.GET("", h.listAll)
	g.POST("", h.create)
	g.PATCH("/:id", h.update)
	g.DELETE("/:id", h.delete)
}

func (h *Handler) listPublished(c *gin.Context) {
	courses, err := h.svc.ListPublished(c.Request.Context())
	if err != nil {
		writeErr(c, err)
		return
	}
	out := make([]Response, len(courses))
	for i, cc := range courses {
		out[i] = ToResponse(&cc)
	}
	response.OK(c, http.StatusOK, out)
}

func (h *Handler) getBySlug(c *gin.Context) {
	// Param tên là "courseId" nhưng giá trị vẫn là slug (không phải UUID) —
	// đặt tên vậy chỉ để tránh conflict wildcard với
	// GET /courses/:courseId/lessons ở lesson module (gin router yêu cầu
	// cùng 1 tên wildcard tại cùng vị trí trong cây route).
	cc, err := h.svc.GetBySlug(c.Request.Context(), c.Param("courseId"))
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToResponse(cc))
}

func (h *Handler) listAll(c *gin.Context) {
	courses, err := h.svc.ListAll(c.Request.Context())
	if err != nil {
		writeErr(c, err)
		return
	}
	out := make([]Response, len(courses))
	for i, cc := range courses {
		out[i] = ToResponse(&cc)
	}
	response.OK(c, http.StatusOK, out)
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}
	cc, err := h.svc.Create(c.Request.Context(), req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusCreated, ToResponse(cc))
}

func (h *Handler) update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid course id")
		return
	}
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}
	cc, err := h.svc.Update(c.Request.Context(), id, req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToResponse(cc))
}

func (h *Handler) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid course id")
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