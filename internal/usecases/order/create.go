package order

import (
	"context"

	"wappi/internal/domain"
	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// CreateInput represents the input for creating an order
type CreateInput struct {
	ProfileID string `json:"profile_id" binding:"required"`
	ETA       string `json:"eta"`
}

// CreateOutput represents the output after creating an order
type CreateOutput struct {
	Data OrderOutputData `json:"data"`
}

// CreateUsecase defines the interface for creating orders
type CreateUsecase interface {
	Execute(ctx context.Context, input CreateInput) (*CreateOutput, apperrors.ApplicationError)
}

type createUsecase struct {
	contextFactory appcontext.Factory
}

// NewCreateUsecase creates a new instance of CreateUsecase
func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{contextFactory: contextFactory}
}

// Execute creates a new order
func (u *createUsecase) Execute(ctx context.Context, input CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	newOrder := &domain.Order{
		ProfileID: &input.ProfileID,
		ETA:       input.ETA,
	}

	created, err := app.Repositories.Order.Create(ctx, newOrder)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toOrderOutputData(created, false),
	}, nil
}
