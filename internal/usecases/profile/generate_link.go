package profile

import (
	"context"
	"time"

	"wappi/internal/domain"
	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// GenerateLinkInput represents the input for generating a profile link
type GenerateLinkInput struct {
	UserID string `json:"user_id" binding:"required"`
}

// GenerateLinkOutput represents the output with the generated link
type GenerateLinkOutput struct {
	Link      string `json:"link"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

// GenerateLinkUsecase defines the interface for generating profile links
type GenerateLinkUsecase interface {
	Execute(ctx context.Context, input GenerateLinkInput) (*GenerateLinkOutput, apperrors.ApplicationError)
}

type generateLinkUsecase struct {
	contextFactory appcontext.Factory
}

// NewGenerateLinkUsecase creates a new instance of GenerateLinkUsecase
func NewGenerateLinkUsecase(contextFactory appcontext.Factory) GenerateLinkUsecase {
	return &generateLinkUsecase{contextFactory: contextFactory}
}

// Execute generates a profile completion link
func (u *generateLinkUsecase) Execute(ctx context.Context, input GenerateLinkInput) (*GenerateLinkOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	// Create token with 24 hour expiration
	expiresAt := time.Now().Add(24 * time.Hour)

	token := &domain.ProfileToken{
		UserID:    input.UserID,
		Used:      false,
		ExpiresAt: expiresAt,
	}

	created, err := app.Repositories.Profile.CreateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	// Generate the frontend link
	link := app.ConfigService.FrontendURL + "/complete-profile/" + created.Token

	return &GenerateLinkOutput{
		Link:      link,
		Token:     created.Token,
		ExpiresAt: expiresAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}
