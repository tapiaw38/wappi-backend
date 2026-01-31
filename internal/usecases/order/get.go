package order

import (
	"context"

	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"

	"github.com/google/uuid"
)

// GetOutput represents the output for getting an order
type GetOutput struct {
	Data OrderOutputData `json:"data"`
}

// GetUsecase defines the interface for getting orders
type GetUsecase interface {
	Execute(ctx context.Context, id string) (*GetOutput, apperrors.ApplicationError)
}

type getUsecase struct {
	contextFactory appcontext.Factory
}

// NewGetUsecase creates a new instance of GetUsecase
func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{contextFactory: contextFactory}
}

// Execute retrieves an order by ID
func (u *getUsecase) Execute(ctx context.Context, id string) (*GetOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	if _, err := uuid.Parse(id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.OrderInvalidIDError, err)
	}

	orderData, err := app.Repositories.Order.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetOutput{
		Data: toOrderOutputData(orderData, true),
	}, nil
}
