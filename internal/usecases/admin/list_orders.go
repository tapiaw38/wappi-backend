package admin

import (
	"context"

	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// ListOrdersOutput represents the output for listing orders
type ListOrdersOutput struct {
	Orders []OrderOutput `json:"orders"`
	Total  int           `json:"total"`
}

// ListOrdersUsecase defines the interface for listing orders
type ListOrdersUsecase interface {
	Execute(ctx context.Context) (*ListOrdersOutput, apperrors.ApplicationError)
}

type listOrdersUsecase struct {
	contextFactory appcontext.Factory
}

// NewListOrdersUsecase creates a new instance of ListOrdersUsecase
func NewListOrdersUsecase(contextFactory appcontext.Factory) ListOrdersUsecase {
	return &listOrdersUsecase{contextFactory: contextFactory}
}

// Execute lists all orders
func (u *listOrdersUsecase) Execute(ctx context.Context) (*ListOrdersOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	orders, err := app.Repositories.Order.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	output := &ListOrdersOutput{
		Orders: make([]OrderOutput, 0, len(orders)),
		Total:  len(orders),
	}

	for _, o := range orders {
		output.Orders = append(output.Orders, toOrderOutput(o))
	}

	return output, nil
}
