package importrecord

import (
	"context"
	"database/sql"

	"yego/internal/domain"
	apperrors "yego/internal/platform/errors"
)

// Repository defines the interface for import record operations
type Repository interface {
	Create(ctx context.Context, record *domain.ImportRecord) (*domain.ImportRecord, apperrors.ApplicationError)
	GetAll(ctx context.Context) ([]*domain.ImportRecord, apperrors.ApplicationError)
	GetByID(ctx context.Context, id string) (*domain.ImportRecord, apperrors.ApplicationError)
	Update(ctx context.Context, id string, data map[string]any, profileID *string) (*domain.ImportRecord, apperrors.ApplicationError)
	Delete(ctx context.Context, id string) apperrors.ApplicationError
	DeleteAll(ctx context.Context) (int64, apperrors.ApplicationError)
}

type repository struct {
	db *sql.DB
}

// NewRepository creates a new import record repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
