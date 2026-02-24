package admin

import (
	"context"

	"yego/internal/domain"
	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	settingsUsecase "yego/internal/usecases/settings"

	"github.com/google/uuid"
)

// UpdateOrderInput represents the input for updating an order
type UpdateOrderInput struct {
	Status        *string           `json:"status,omitempty"`
	StatusMessage *string           `json:"status_message,omitempty"`
	ETA           *string           `json:"eta,omitempty"`
	Data          *domain.OrderData `json:"data,omitempty"`
	Token         string            `json:"-"`
}

// UpdateOrderUsecase defines the interface for updating orders
type UpdateOrderUsecase interface {
	Execute(ctx context.Context, id string, input UpdateOrderInput) (*OrderOutput, apperrors.ApplicationError)
}

type updateOrderUsecase struct {
	contextFactory          appcontext.Factory
	calculateDeliveryFeeUse settingsUsecase.CalculateDeliveryFeeUsecase
}

// NewUpdateOrderUsecase creates a new instance of UpdateOrderUsecase
func NewUpdateOrderUsecase(contextFactory appcontext.Factory, calculateDeliveryFeeUse settingsUsecase.CalculateDeliveryFeeUsecase) UpdateOrderUsecase {
	return &updateOrderUsecase{
		contextFactory:          contextFactory,
		calculateDeliveryFeeUse: calculateDeliveryFeeUse,
	}
}

// Execute updates an order
func (u *updateOrderUsecase) Execute(ctx context.Context, id string, input UpdateOrderInput) (*OrderOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	// Validate order ID
	if _, err := uuid.Parse(id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.OrderInvalidIDError, err)
	}

	// Get existing order (capture status before update)
	order, err := app.Repositories.Order.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	previousStatus := order.Status

	// Update fields if provided
	if input.Status != nil {
		if !domain.IsValidStatus(*input.Status) {
			return nil, apperrors.NewApplicationError(mappings.OrderInvalidStatusError, nil)
		}
		order.Status = domain.OrderStatus(*input.Status)
	}

	if input.StatusMessage != nil {
		order.StatusMessage = input.StatusMessage
	}

	if input.ETA != nil {
		order.ETA = *input.ETA
	}

	if input.Data != nil {
		order.Data = input.Data
	}

	// Save changes
	updatedOrder, err := app.Repositories.Order.Update(ctx, order)
	if err != nil {
		return nil, err
	}

	// Payment processing removed - payments are now processed at order creation
	// Keeping this comment for reference
	_ = previousStatus // Keep variable to avoid compilation error

	output := toOrderOutput(updatedOrder)
	return &output, nil
}
