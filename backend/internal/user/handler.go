package user

import (
	"net/http"
	"strconv"

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

// RegisterRoutes mounts /users/me endpoints on an already-authenticated group.
func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	g := rg.Group("/users")
	g.GET("/me", h.getMe)
	g.PATCH("/me", h.updateMe)
}

// RegisterAdminRoutes mounts admin-only user management endpoints. The caller
// (router.go) is responsible for wrapping rg with
// middleware.RequireAuth + middleware.RequireRole("admin") — this function
// does not add any authorization itself.
func RegisterAdminRoutes(rg *gin.RouterGroup, h *Handler) {
	g := rg.Group("/admin/users")
	g.GET("", h.list)
	g.PATCH("/:id/role", h.updateRole)
}

func (h *Handler) getMe(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	u, err := h.svc.GetByID(c.Request.Context(), userID)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToPublicResponse(u))
}

func (h *Handler) updateMe(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	u, err := h.svc.UpdateProfile(c.Request.Context(), userID, req)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToPublicResponse(u))
}

func (h *Handler) updateRole(c *gin.Context) {
	actorID := c.MustGet("user_id").(uuid.UUID)

	targetID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", "invalid user id")
		return
	}

	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	u, err := h.svc.UpdateRole(c.Request.Context(), actorID, targetID, req.Role)
	if err != nil {
		writeErr(c, err)
		return
	}
	response.OK(c, http.StatusOK, ToPublicResponse(u))
}

func (h *Handler) list(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	users, total, err := h.svc.List(c.Request.Context(), page, pageSize)
	if err != nil {
		writeErr(c, err)
		return
	}

	out := make([]PublicResponse, len(users))
	for i, u := range users {
		out[i] = ToPublicResponse(&u)
	}
	response.OKWithMeta(c, http.StatusOK, out, gin.H{
		"page": page, "page_size": pageSize, "total": total,
	})
}

func writeErr(c *gin.Context, err error) {
	if ae, ok := apperrors.As(err); ok {
		response.Err(c, ae.Status, ae.Code, ae.Message)
		return
	}
	response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
}
