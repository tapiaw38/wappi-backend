package admin

import (
	"wappi/internal/platform/appcontext"
)

// Usecases aggregates all admin-related use cases
type Usecases struct {
	ListProfiles ListProfilesUsecase
	ListOrders   ListOrdersUsecase
	UpdateOrder  UpdateOrderUsecase
}

// NewUsecases creates all admin use cases
func NewUsecases(contextFactory appcontext.Factory) *Usecases {
	return &Usecases{
		ListProfiles: NewListProfilesUsecase(contextFactory),
		ListOrders:   NewListOrdersUsecase(contextFactory),
		UpdateOrder:  NewUpdateOrderUsecase(contextFactory),
	}
}
