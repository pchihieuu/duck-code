// dto.go
package lesson

import "github.com/google/uuid"

type CreateRequest struct {
	CourseID   uuid.UUID `json:"course_id" binding:"required"`
	Title      string    `json:"title" binding:"required,min=2,max=150"`
	Slug       string    `json:"slug" binding:"required,min=2,max=150"`
	ContentMDX string    `json:"content_mdx" binding:"omitempty"`
	OrderIndex int       `json:"order_index"`
}

type UpdateRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=2,max=150"`
	ContentMDX  *string `json:"content_mdx"`
	OrderIndex  *int    `json:"order_index"`
	IsPublished *bool   `json:"is_published"`
}

type ListItemResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	OrderIndex  int       `json:"order_index"`
	IsPublished bool      `json:"is_published"`
}

type DetailResponse struct {
	ListItemResponse
	ContentMDX string `json:"content_mdx"`
}

func ToListItem(l *Lesson) ListItemResponse {
	return ListItemResponse{
		ID: l.ID, Title: l.Title, Slug: l.Slug,
		OrderIndex: l.OrderIndex, IsPublished: l.IsPublished,
	}
}

func ToDetail(l *Lesson) DetailResponse {
	return DetailResponse{ListItemResponse: ToListItem(l), ContentMDX: l.ContentMDX}
}