package profile

import (
	"yego/internal/domain"
)

// ProfileOutputData represents basic profile data for outputs
type ProfileOutputData struct {
	ID          string          `json:"id"`
	UserID      string          `json:"user_id"`
	PhoneNumber string          `json:"phone_number"`
	Location    *LocationOutput `json:"location,omitempty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

// LocationOutput represents location data in the output
type LocationOutput struct {
	ID        string  `json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
}

// toProfileOutputData converts a domain profile to output data
func toProfileOutputData(profile *domain.Profile, location *domain.ProfileLocation) ProfileOutputData {
	output := ProfileOutputData{
		ID:          profile.ID,
		UserID:      profile.UserID,
		PhoneNumber: profile.PhoneNumber,
		CreatedAt:   profile.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   profile.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	if location != nil {
		output.Location = &LocationOutput{
			ID:        location.ID,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
			Address:   location.Address,
		}
	}

	return output
}
