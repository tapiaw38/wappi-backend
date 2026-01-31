package order

import (
	"wappi/internal/platform/appcontext"
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
func NewUsecases(contextFactory appcontext.Factory) *Usecases {
	return &Usecases{
		Create:         NewCreateUsecase(contextFactory),
		CreateWithLink: NewCreateWithLinkUsecase(contextFactory),
		Claim:          NewClaimUsecase(contextFactory),
		Get:            NewGetUsecase(contextFactory),
		UpdateStatus:   NewUpdateStatusUsecase(contextFactory),
		ListMyOrders:   NewListMyOrdersUsecase(contextFactory),
	}
}
