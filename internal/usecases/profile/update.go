package profile

import (
	"context"

	"yego/internal/domain"
	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"

	"github.com/google/uuid"
)

// UpdateProfileInput represents the input for updating a profile
type UpdateProfileInput struct {
	PhoneNumber string  `json:"phone_number"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

// UpdateProfileUsecase defines the interface for updating profiles
type UpdateProfileUsecase interface {
	Execute(ctx context.Context, id string, input UpdateProfileInput) (*GetProfileOutput, apperrors.ApplicationError)
}

type updateProfileUsecase struct {
	contextFactory appcontext.Factory
}

// NewUpdateProfileUsecase creates a new instance of UpdateProfileUsecase
func NewUpdateProfileUsecase(contextFactory appcontext.Factory) UpdateProfileUsecase {
	return &updateProfileUsecase{contextFactory: contextFactory}
}

// Execute updates a profile
func (u *updateProfileUsecase) Execute(ctx context.Context, id string, input UpdateProfileInput) (*GetProfileOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	// Validate profile ID
	if _, err := uuid.Parse(id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.InvalidUserIDError, err)
	}

	// Get existing profile
	existingProfile, err := app.Repositories.Profile.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Create or update location
	location := &domain.ProfileLocation{
		ID:        uuid.New().String(),
		Longitude: input.Longitude,
		Latitude:  input.Latitude,
		Address:   input.Address,
	}

	savedLocation, locErr := app.Repositories.Profile.CreateLocation(ctx, location)
	if locErr != nil {
		return nil, locErr
	}

	// Update profile
	existingProfile.PhoneNumber = input.PhoneNumber
	existingProfile.LocationID = &savedLocation.ID

	updatedProfile, err := app.Repositories.Profile.Update(ctx, existingProfile)
	if err != nil {
		return nil, err
	}

	return &GetProfileOutput{
		Data: toProfileOutputData(updatedProfile, savedLocation),
	}, nil
}
