package order

import (
	"wappi/internal/platform/appcontext"
	"wappi/internal/usecases/notification"
)

// Usecases aggregates all order-related use cases
type Usecases struct {
	Create         CreateUsecase
	CreateWithLink CreateWithLinkUsecase
	Claim          ClaimUsecase
	Get            GetUsecase
	UpdateStatus   UpdateStatusUsecase
	ListMyOrders   ListMyOrdersUsecase
}

// NewUsecases creates all order use cases
func NewUsecases(contextFactory appcontext.Factory, notificationSvc notification.Service) *Usecases {
	return &Usecases{
		Create:         NewCreateUsecase(contextFactory),
		CreateWithLink: NewCreateWithLinkUsecase(contextFactory),
		Claim:          NewClaimUsecase(contextFactory, notificationSvc),
		Get:            NewGetUsecase(contextFactory),
		UpdateStatus:   NewUpdateStatusUsecase(contextFactory),
		ListMyOrders:   NewListMyOrdersUsecase(contextFactory),
	}
}
