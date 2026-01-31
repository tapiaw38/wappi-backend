package profile

import (
	"context"

	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// ValidateTokenOutput represents the output for token validation
type ValidateTokenOutput struct {
	Valid   bool   `json:"valid"`
	UserID  string `json:"user_id"`
	Message string `json:"message,omitempty"`
}

// ValidateTokenUsecase defines the interface for validating profile tokens
type ValidateTokenUsecase interface {
	Execute(ctx context.Context, token string) (*ValidateTokenOutput, apperrors.ApplicationError)
}

type validateTokenUsecase struct {
	contextFactory appcontext.Factory
}

// NewValidateTokenUsecase creates a new instance of ValidateTokenUsecase
func NewValidateTokenUsecase(contextFactory appcontext.Factory) ValidateTokenUsecase {
	return &validateTokenUsecase{contextFactory: contextFactory}
}

// Execute validates a profile token
func (u *validateTokenUsecase) Execute(ctx context.Context, token string) (*ValidateTokenOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	profileToken, err := app.Repositories.Profile.GetToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &ValidateTokenOutput{
		Valid:  true,
		UserID: profileToken.UserID,
	}, nil
}
