package admin

import (
	"context"

	"wappi/internal/platform/appcontext"
	apperrors "wappi/internal/platform/errors"
)

// ListProfilesOutput represents the output for listing profiles
type ListProfilesOutput struct {
	Profiles []ProfileOutput `json:"profiles"`
	Total    int             `json:"total"`
}

// ListProfilesUsecase defines the interface for listing profiles
type ListProfilesUsecase interface {
	Execute(ctx context.Context) (*ListProfilesOutput, apperrors.ApplicationError)
}

type listProfilesUsecase struct {
	contextFactory appcontext.Factory
}

// NewListProfilesUsecase creates a new instance of ListProfilesUsecase
func NewListProfilesUsecase(contextFactory appcontext.Factory) ListProfilesUsecase {
	return &listProfilesUsecase{contextFactory: contextFactory}
}

// Execute lists all profiles
func (u *listProfilesUsecase) Execute(ctx context.Context) (*ListProfilesOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	profiles, err := app.Repositories.Profile.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	output := &ListProfilesOutput{
		Profiles: make([]ProfileOutput, 0, len(profiles)),
		Total:    len(profiles),
	}

	for _, p := range profiles {
		profileOutput := ProfileOutput{
			ID:          p.ID,
			UserID:      p.UserID,
			PhoneNumber: p.PhoneNumber,
			CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   p.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}

		// Get location if exists
		if p.LocationID != nil {
			location, locErr := app.Repositories.Profile.GetLocationByID(ctx, *p.LocationID)
			if locErr == nil && location != nil {
				profileOutput.Location = &LocationOutput{
					ID:        location.ID,
					Longitude: location.Longitude,
					Latitude:  location.Latitude,
					Address:   location.Address,
				}
			}
		}

		output.Profiles = append(output.Profiles, profileOutput)
	}

	return output, nil
}
