package profile

import (
	"context"

	"wappi/internal/domain"
	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// CompleteProfileInput represents the input for completing a profile
type CompleteProfileInput struct {
	Token       string  `json:"token" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Address     string  `json:"address"`
}

// CompleteProfileOutput represents the output after completing a profile
type CompleteProfileOutput struct {
	Data ProfileOutputData `json:"data"`
}

// CompleteProfileUsecase defines the interface for completing profiles
type CompleteProfileUsecase interface {
	Execute(ctx context.Context, input CompleteProfileInput) (*CompleteProfileOutput, apperrors.ApplicationError)
}

type completeProfileUsecase struct {
	contextFactory appcontext.Factory
}

// NewCompleteProfileUsecase creates a new instance of CompleteProfileUsecase
func NewCompleteProfileUsecase(contextFactory appcontext.Factory) CompleteProfileUsecase {
	return &completeProfileUsecase{contextFactory: contextFactory}
}

// Execute completes a user profile
func (u *completeProfileUsecase) Execute(ctx context.Context, input CompleteProfileInput) (*CompleteProfileOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	// Validate token
	profileToken, err := app.Repositories.Profile.GetToken(ctx, input.Token)
	if err != nil {
		return nil, err
	}

	// Create location first
	location := &domain.ProfileLocation{
		Longitude: input.Longitude,
		Latitude:  input.Latitude,
		Address:   input.Address,
	}

	createdLocation, err := app.Repositories.Profile.CreateLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	// Check if profile already exists for this user
	existingProfile, _ := app.Repositories.Profile.GetByUserID(ctx, profileToken.UserID)

	var resultProfile *domain.Profile

	if existingProfile != nil {
		// Update existing profile
		existingProfile.PhoneNumber = input.PhoneNumber
		existingProfile.LocationID = &createdLocation.ID
		resultProfile, err = app.Repositories.Profile.Update(ctx, existingProfile)
		if err != nil {
			return nil, err
		}
	} else {
		// Create new profile
		newProfile := &domain.Profile{
			UserID:      profileToken.UserID,
			PhoneNumber: input.PhoneNumber,
			LocationID:  &createdLocation.ID,
		}
		resultProfile, err = app.Repositories.Profile.Create(ctx, newProfile)
		if err != nil {
			return nil, err
		}
	}

	// Mark token as used
	if markErr := app.Repositories.Profile.MarkTokenUsed(ctx, input.Token); markErr != nil {
		return nil, markErr
	}

	return &CompleteProfileOutput{
		Data: toProfileOutputData(resultProfile, createdLocation),
	}, nil
}
