package order

import (
	"context"

	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
)

// GetClaimInfoInput represents the input for getting claim info
type GetClaimInfoInput struct {
	Token string `json:"token" binding:"required"`
}

// GetClaimInfoOutput represents the output for claim info
type GetClaimInfoOutput struct {
	OrderID   string  `json:"order_id"`
	UserID    *string `json:"user_id"`
	ProfileID *string `json:"profile_id"`
	Status    string  `json:"status"`
	IsClaimed bool    `json:"is_claimed"`
}

// GetClaimInfoUsecase defines the interface for getting claim info
type GetClaimInfoUsecase interface {
	Execute(ctx context.Context, input GetClaimInfoInput) (*GetClaimInfoOutput, apperrors.ApplicationError)
}

type getClaimInfoUsecase struct {
	contextFactory appcontext.Factory
}

// NewGetClaimInfoUsecase creates a new instance of GetClaimInfoUsecase
func NewGetClaimInfoUsecase(contextFactory appcontext.Factory) GetClaimInfoUsecase {
	return &getClaimInfoUsecase{contextFactory: contextFactory}
}

// Execute gets order info from a claim token without claiming it
func (u *getClaimInfoUsecase) Execute(ctx context.Context, input GetClaimInfoInput) (*GetClaimInfoOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	// Get the order token
	orderToken, err := app.Repositories.OrderToken.GetByToken(ctx, input.Token)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.OrderTokenNotFoundError, err)
	}

	// Get the order
	order, err := app.Repositories.Order.GetByID(ctx, orderToken.OrderID)
	if err != nil {
		return nil, err
	}

	return &GetClaimInfoOutput{
		OrderID:   order.ID,
		UserID:    order.UserID,
		ProfileID: order.ProfileID,
		Status:    string(order.Status),
		IsClaimed: order.UserID != nil && *order.UserID != "",
	}, nil
}
