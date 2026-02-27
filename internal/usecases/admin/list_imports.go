package admin

import (
	"context"
	"time"

	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
)

// ImportRecordOutput represents an import record in API responses
type ImportRecordOutput struct {
	ID        string         `json:"id"`
	Data      map[string]any `json:"data,omitempty"`
	ProfileID *string        `json:"profile_id,omitempty"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

// ListImportsOutput is the result of listing import records
type ListImportsOutput struct {
	Records []ImportRecordOutput `json:"records"`
	Total   int                  `json:"total"`
}

// ListImportsUsecase defines the interface for listing import records
type ListImportsUsecase interface {
	Execute(ctx context.Context) (*ListImportsOutput, apperrors.ApplicationError)
}

type listImportsUsecase struct {
	contextFactory appcontext.Factory
}

// NewListImportsUsecase creates a new instance of ListImportsUsecase
func NewListImportsUsecase(contextFactory appcontext.Factory) ListImportsUsecase {
	return &listImportsUsecase{contextFactory: contextFactory}
}

// Execute retrieves all import records
func (u *listImportsUsecase) Execute(ctx context.Context) (*ListImportsOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	records, err := app.Repositories.ImportRecord.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	output := &ListImportsOutput{
		Records: make([]ImportRecordOutput, 0, len(records)),
		Total:   len(records),
	}

	for _, r := range records {
		output.Records = append(output.Records, ImportRecordOutput{
			ID:        r.ID,
			Data:      r.Data,
			ProfileID: r.ProfileID,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
			UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
		})
	}

	return output, nil
}
