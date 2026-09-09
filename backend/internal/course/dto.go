// dto.go
package course

import "github.com/google/uuid"

type CreateRequest struct {
	Title       string `json:"title" binding:"required,min=2,max=150"`
	Slug        string `json:"slug" binding:"required,min=2,max=150"`
	Description string `json:"description"`
	Language    string `json:"language" binding:"required"`
	OrderIndex  int    `json:"order_index"`
}

type UpdateRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=2,max=150"`
	Description *string `json:"description"`
	Language    *string `json:"language"`
	OrderIndex  *int    `json:"order_index"`
	IsPublished *bool   `json:"is_published"`
}

type Response struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	OrderIndex  int       `json:"order_index"`
	IsPublished bool      `json:"is_published"`
}

func ToResponse(c *Course) Response {
	return Response{
		ID: c.ID, Title: c.Title, Slug: c.Slug, Description: c.Description,
		Language: c.Language, OrderIndex: c.OrderIndex, IsPublished: c.IsPublished,
	}
}