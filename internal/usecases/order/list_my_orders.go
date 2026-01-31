package order

import (
	"context"

	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// ListMyOrdersOutput represents the output for listing user's orders
type ListMyOrdersOutput struct {
	Orders []OrderOutputData `json:"orders"`
	Total  int               `json:"total"`
}

// ListMyOrdersUsecase defines the interface for listing user's orders
type ListMyOrdersUsecase interface {
	Execute(ctx context.Context, userID string) (*ListMyOrdersOutput, apperrors.ApplicationError)
}

type listMyOrdersUsecase struct {
	contextFactory appcontext.Factory
}

// NewListMyOrdersUsecase creates a new instance of ListMyOrdersUsecase
func NewListMyOrdersUsecase(contextFactory appcontext.Factory) ListMyOrdersUsecase {
	return &listMyOrdersUsecase{contextFactory: contextFactory}
}

// Execute lists all orders for a specific user
func (u *listMyOrdersUsecase) Execute(ctx context.Context, userID string) (*ListMyOrdersOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	orders, err := app.Repositories.Order.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	output := &ListMyOrdersOutput{
		Orders: make([]OrderOutputData, 0, len(orders)),
		Total:  len(orders),
	}

	for _, o := range orders {
		output.Orders = append(output.Orders, toOrderOutputData(o, true))
	}

	return output, nil
}
