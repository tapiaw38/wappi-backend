package usecases

import (
	"wappi/internal/platform/appcontext"
	"wappi/internal/usecases/admin"
	"wappi/internal/usecases/order"
	"wappi/internal/usecases/profile"
	"wappi/internal/usecases/settings"
)

type Usecases struct {
	Order    Order
	Profile  Profile
	Admin    Admin
	Settings Settings
}

type Order struct {
	CreateUsecase         order.CreateUsecase
	CreateWithLinkUsecase order.CreateWithLinkUsecase
	ClaimUsecase          order.ClaimUsecase
	GetUsecase            order.GetUsecase
	UpdateStatusUsecase   order.UpdateStatusUsecase
	ListMyOrdersUsecase   order.ListMyOrdersUsecase
}

type Profile struct {
	GenerateLinkUsecase    profile.GenerateLinkUsecase
	ValidateTokenUsecase   profile.ValidateTokenUsecase
	CompleteProfileUsecase profile.CompleteProfileUsecase
	GetUsecase             profile.GetProfileUsecase
	UpdateUsecase          profile.UpdateProfileUsecase
	CheckCompletedUsecase  profile.CheckCompletedUsecase
}

type Admin struct {
	ListProfilesUsecase admin.ListProfilesUsecase
	ListOrdersUsecase   admin.ListOrdersUsecase
	UpdateOrderUsecase  admin.UpdateOrderUsecase
}

type Settings struct {
	GetUsecase                  settings.GetUsecase
	UpdateUsecase               settings.UpdateUsecase
	CalculateDeliveryFeeUsecase settings.CalculateDeliveryFeeUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *Usecases {
	return &Usecases{
		Order: Order{
			CreateUsecase:         order.NewCreateUsecase(contextFactory),
			CreateWithLinkUsecase: order.NewCreateWithLinkUsecase(contextFactory),
			ClaimUsecase:          order.NewClaimUsecase(contextFactory),
			GetUsecase:            order.NewGetUsecase(contextFactory),
			UpdateStatusUsecase:   order.NewUpdateStatusUsecase(contextFactory),
			ListMyOrdersUsecase:   order.NewListMyOrdersUsecase(contextFactory),
		},
		Profile: Profile{
			GenerateLinkUsecase:    profile.NewGenerateLinkUsecase(contextFactory),
			ValidateTokenUsecase:   profile.NewValidateTokenUsecase(contextFactory),
			CompleteProfileUsecase: profile.NewCompleteProfileUsecase(contextFactory),
			GetUsecase:             profile.NewGetProfileUsecase(contextFactory),
			UpdateUsecase:          profile.NewUpdateProfileUsecase(contextFactory),
			CheckCompletedUsecase:  profile.NewCheckCompletedUsecase(contextFactory),
		},
		Admin: Admin{
			ListProfilesUsecase: admin.NewListProfilesUsecase(contextFactory),
			ListOrdersUsecase:   admin.NewListOrdersUsecase(contextFactory),
			UpdateOrderUsecase:  admin.NewUpdateOrderUsecase(contextFactory),
		},
		Settings: Settings{
			GetUsecase:                  settings.NewGetUsecase(contextFactory),
			UpdateUsecase:               settings.NewUpdateUsecase(contextFactory),
			CalculateDeliveryFeeUsecase: settings.NewCalculateDeliveryFeeUsecase(contextFactory),
		},
	}
}
