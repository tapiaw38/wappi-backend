package order

import (
	"context"
	"fmt"
	"log"

	"wappi/internal/domain"
	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	settingsUsecase "wappi/internal/usecases/settings"
)

// CreateInput represents the input for creating an order
type CreateInput struct {
	ProfileID    string `json:"profile_id" binding:"required"`
	ETA          string `json:"eta"`
	SecurityCode string `json:"security_code"`
	Token        string
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
	contextFactory          appcontext.Factory
	calculateDeliveryFeeUse settingsUsecase.CalculateDeliveryFeeUsecase
}

// NewCreateUsecase creates a new instance of CreateUsecase
func NewCreateUsecase(contextFactory appcontext.Factory, calculateDeliveryFeeUse settingsUsecase.CalculateDeliveryFeeUsecase) CreateUsecase {
	return &createUsecase{
		contextFactory:          contextFactory,
		calculateDeliveryFeeUse: calculateDeliveryFeeUse,
	}
}

// Execute creates a new order and processes payment immediately
func (u *createUsecase) Execute(ctx context.Context, input CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	newOrder := &domain.Order{
		ProfileID: &input.ProfileID,
		ETA:       input.ETA,
		Status:    domain.StatusCreated,
	}

	created, err := app.Repositories.Order.Create(ctx, newOrder)
	if err != nil {
		return nil, err
	}

	// Process payment immediately if security code is provided
	if input.SecurityCode != "" {
		paymentErr := ProcessPaymentForOrder(ctx, app, created, input.Token, input.SecurityCode, u.calculateDeliveryFeeUse)
		if paymentErr != nil {
			log.Printf("Payment failed for order %s: %v", created.ID, paymentErr)
			// Keep status as CREATED but return error
			return nil, apperrors.NewApplicationError(mappings.OrderPaymentFailedError, fmt.Errorf("payment failed: %w", paymentErr))
		}
		// Payment successful - update status to CONFIRMED
		created.Status = domain.StatusConfirmed
		created, _ = app.Repositories.Order.Update(ctx, created)
	}

	return &CreateOutput{
		Data: toOrderOutputData(created, false),
	}, nil
}
