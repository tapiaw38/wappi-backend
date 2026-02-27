package admin

import (
	"context"

	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
)

// DeleteImportUsecase defines the interface for deleting an import record
type DeleteImportUsecase interface {
	Execute(ctx context.Context, id string) apperrors.ApplicationError
}

type deleteImportUsecase struct {
	contextFactory appcontext.Factory
}

// NewDeleteImportUsecase creates a new instance of DeleteImportUsecase
func NewDeleteImportUsecase(contextFactory appcontext.Factory) DeleteImportUsecase {
	return &deleteImportUsecase{contextFactory: contextFactory}
}

// Execute deletes an import record by ID
func (u *deleteImportUsecase) Execute(ctx context.Context, id string) apperrors.ApplicationError {
	app := u.contextFactory()
	return app.Repositories.ImportRecord.Delete(ctx, id)
}
