package admin

import (
	"context"
	"time"

	"yego/internal/domain"
	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
)

// CreateImportInput holds the data for a new import record
type CreateImportInput struct {
	Data      map[string]any `json:"data"`
	ProfileID *string        `json:"profile_id"`
}

// CreateImportUsecase defines the interface for creating a single import record
type CreateImportUsecase interface {
	Execute(ctx context.Context, input CreateImportInput) (*ImportRecordOutput, apperrors.ApplicationError)
}

type createImportUsecase struct {
	contextFactory appcontext.Factory
}

// NewCreateImportUsecase creates a new instance of CreateImportUsecase
func NewCreateImportUsecase(contextFactory appcontext.Factory) CreateImportUsecase {
	return &createImportUsecase{contextFactory: contextFactory}
}

// Execute creates a single import record
func (u *createImportUsecase) Execute(ctx context.Context, input CreateImportInput) (*ImportRecordOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	record := &domain.ImportRecord{
		Data:      input.Data,
		ProfileID: input.ProfileID,
	}

	created, err := app.Repositories.ImportRecord.Create(ctx, record)
	if err != nil {
		return nil, err
	}

	return &ImportRecordOutput{
		ID:        created.ID,
		Data:      created.Data,
		ProfileID: created.ProfileID,
		CreatedAt: created.CreatedAt.Format(time.RFC3339),
		UpdatedAt: created.UpdatedAt.Format(time.RFC3339),
	}, nil
}
