package profile

import (
	"context"

	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"

	"github.com/google/uuid"
)

// GetProfileOutput represents the output for getting a profile
type GetProfileOutput struct {
	Data ProfileOutputData `json:"data"`
}

// GetProfileUsecase defines the interface for getting profiles
type GetProfileUsecase interface {
	Execute(ctx context.Context, id string) (*GetProfileOutput, apperrors.ApplicationError)
}

type getProfileUsecase struct {
	contextFactory appcontext.Factory
}

// NewGetProfileUsecase creates a new instance of GetProfileUsecase
func NewGetProfileUsecase(contextFactory appcontext.Factory) GetProfileUsecase {
	return &getProfileUsecase{contextFactory: contextFactory}
}

// Execute retrieves a profile by ID
func (u *getProfileUsecase) Execute(ctx context.Context, id string) (*GetProfileOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	if _, err := uuid.Parse(id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.InvalidUserIDError, err)
	}

	profile, err := app.Repositories.Profile.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Get location if exists
	var location *LocationOutput
	if profile.LocationID != nil {
		loc, locErr := app.Repositories.Profile.GetLocationByID(ctx, *profile.LocationID)
		if locErr == nil && loc != nil {
			location = &LocationOutput{
				ID:        loc.ID,
				Longitude: loc.Longitude,
				Latitude:  loc.Latitude,
				Address:   loc.Address,
			}
		}
	}

	output := ProfileOutputData{
		ID:          profile.ID,
		UserID:      profile.UserID,
		PhoneNumber: profile.PhoneNumber,
		Location:    location,
		CreatedAt:   profile.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   profile.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return &GetProfileOutput{
		Data: output,
	}, nil
}
