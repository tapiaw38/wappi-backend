package admin

import (
	"context"
	"time"

	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
)

// UpdateImportInput holds the fields that can be updated on an import record
type UpdateImportInput struct {
	Data      map[string]any `json:"data"`
	ProfileID *string        `json:"profile_id"`
}

// UpdateImportUsecase defines the interface for updating an import record
type UpdateImportUsecase interface {
	Execute(ctx context.Context, id string, input UpdateImportInput) (*ImportRecordOutput, apperrors.ApplicationError)
}

type updateImportUsecase struct {
	contextFactory appcontext.Factory
}

// NewUpdateImportUsecase creates a new instance of UpdateImportUsecase
func NewUpdateImportUsecase(contextFactory appcontext.Factory) UpdateImportUsecase {
	return &updateImportUsecase{contextFactory: contextFactory}
}

// Execute updates data and profile_id on an import record
func (u *updateImportUsecase) Execute(ctx context.Context, id string, input UpdateImportInput) (*ImportRecordOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	record, err := app.Repositories.ImportRecord.Update(ctx, id, input.Data, input.ProfileID)
	if err != nil {
		return nil, err
	}

	return &ImportRecordOutput{
		ID:        record.ID,
		Data:      record.Data,
		ProfileID: record.ProfileID,
		CreatedAt: record.CreatedAt.Format(time.RFC3339),
		UpdatedAt: record.UpdatedAt.Format(time.RFC3339),
	}, nil
}
