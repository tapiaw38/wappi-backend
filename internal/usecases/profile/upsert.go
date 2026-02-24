package profile

import (
	"context"

	"yego/internal/domain"
	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"

	"github.com/google/uuid"
)

type UpsertProfileInput struct {
	PhoneNumber string  `json:"phone_number"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

type UpsertProfileUsecase interface {
	Execute(ctx context.Context, userID string, input UpsertProfileInput) (*GetProfileOutput, apperrors.ApplicationError)
}

type upsertProfileUsecase struct {
	contextFactory appcontext.Factory
}

func NewUpsertProfileUsecase(contextFactory appcontext.Factory) UpsertProfileUsecase {
	return &upsertProfileUsecase{contextFactory: contextFactory}
}

func (u *upsertProfileUsecase) Execute(ctx context.Context, userID string, input UpsertProfileInput) (*GetProfileOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	existingProfile, _ := app.Repositories.Profile.GetByUserID(ctx, userID)

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

	var resultProfile *domain.Profile
	var err apperrors.ApplicationError

	if existingProfile != nil {
		existingProfile.PhoneNumber = input.PhoneNumber
		existingProfile.LocationID = &savedLocation.ID
		resultProfile, err = app.Repositories.Profile.Update(ctx, existingProfile)
		if err != nil {
			return nil, err
		}
	} else {
		newProfile := &domain.Profile{
			UserID:      userID,
			PhoneNumber: input.PhoneNumber,
			LocationID:  &savedLocation.ID,
		}
		resultProfile, err = app.Repositories.Profile.Create(ctx, newProfile)
		if err != nil {
			return nil, err
		}
	}

	return &GetProfileOutput{
		Data: toProfileOutputData(resultProfile, savedLocation),
	}, nil
}
