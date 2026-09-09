// repository.go
package course

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("course not found")

type Repository interface {
	Create(ctx context.Context, c *Course) error
	FindByID(ctx context.Context, id uuid.UUID) (*Course, error)
	FindBySlug(ctx context.Context, slug string) (*Course, error)
	ListPublished(ctx context.Context) ([]Course, error)
	ListAll(ctx context.Context) ([]Course, error)
	Update(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	SoftDelete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, c *Course) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*Course, error) {
	return r.findOne(ctx, "id = ?", id)
}

func (r *repository) FindBySlug(ctx context.Context, slug string) (*Course, error) {
	return r.findOne(ctx, "slug = ?", slug)
}

func (r *repository) ListPublished(ctx context.Context) ([]Course, error) {
	var courses []Course
	err := r.db.WithContext(ctx).Where("is_published = ?", true).Order("order_index ASC").Find(&courses).Error
	return courses, err
}

func (r *repository) ListAll(ctx context.Context) ([]Course, error) {
	var courses []Course
	err := r.db.WithContext(ctx).Order("order_index ASC").Find(&courses).Error
	return courses, err
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&Course{}).Where("id = ?", id).Updates(updates).Error
}

func (r *repository) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Course{}, "id = ?", id).Error
}

func (r *repository) findOne(ctx context.Context, cond string, args ...interface{}) (*Course, error) {
	var c Course
	err := r.db.WithContext(ctx).Where(cond, args...).First(&c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &c, nil
}