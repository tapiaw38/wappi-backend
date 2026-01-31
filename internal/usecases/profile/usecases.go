package profile

import (
	"wappi/internal/platform/appcontext"
)

// Usecases aggregates all profile-related use cases
type Usecases struct {
	GenerateLink    GenerateLinkUsecase
	ValidateToken   ValidateTokenUsecase
	CompleteProfile CompleteProfileUsecase
	Get             GetProfileUsecase
	Update          UpdateProfileUsecase
	CheckCompleted  CheckCompletedUsecase
}

// NewUsecases creates all profile use cases
func NewUsecases(contextFactory appcontext.Factory) *Usecases {
	return &Usecases{
		GenerateLink:    NewGenerateLinkUsecase(contextFactory),
		ValidateToken:   NewValidateTokenUsecase(contextFactory),
		CompleteProfile: NewCompleteProfileUsecase(contextFactory),
		Get:             NewGetProfileUsecase(contextFactory),
		Update:          NewUpdateProfileUsecase(contextFactory),
		CheckCompleted:  NewCheckCompletedUsecase(contextFactory),
	}
}
