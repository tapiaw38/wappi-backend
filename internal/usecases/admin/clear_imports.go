package admin

import (
	"context"

	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
)

// ClearImportsOutput is the result of clearing all import records
type ClearImportsOutput struct {
	Deleted int64 `json:"deleted"`
}

// ClearImportsUsecase defines the interface for deleting all import records
type ClearImportsUsecase interface {
	Execute(ctx context.Context) (*ClearImportsOutput, apperrors.ApplicationError)
}

type clearImportsUsecase struct {
	contextFactory appcontext.Factory
}

// NewClearImportsUsecase creates a new instance of ClearImportsUsecase
func NewClearImportsUsecase(contextFactory appcontext.Factory) ClearImportsUsecase {
	return &clearImportsUsecase{contextFactory: contextFactory}
}

// Execute deletes all import records and returns the count deleted
func (u *clearImportsUsecase) Execute(ctx context.Context) (*ClearImportsOutput, apperrors.ApplicationError) {
	app := u.contextFactory()
	deleted, appErr := app.Repositories.ImportRecord.DeleteAll(ctx)
	if appErr != nil {
		return nil, appErr
	}
	return &ClearImportsOutput{Deleted: deleted}, nil
}
