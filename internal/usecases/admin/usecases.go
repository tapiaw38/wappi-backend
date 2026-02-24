package admin

import (
	"yego/internal/platform/appcontext"
	settingsUsecase "yego/internal/usecases/settings"
)

// Usecases aggregates all admin-related use cases
type Usecases struct {
	ListProfiles     ListProfilesUsecase
	ListOrders       ListOrdersUsecase
	ListTransactions ListTransactionsUsecase
	UpdateOrder      UpdateOrderUsecase
}

// NewUsecases creates all admin use cases
func NewUsecases(contextFactory appcontext.Factory, calculateDeliveryFeeUse settingsUsecase.CalculateDeliveryFeeUsecase) *Usecases {
	return &Usecases{
		ListProfiles:     NewListProfilesUsecase(contextFactory),
		ListOrders:       NewListOrdersUsecase(contextFactory),
		ListTransactions: NewListTransactionsUsecase(contextFactory),
		UpdateOrder:      NewUpdateOrderUsecase(contextFactory, calculateDeliveryFeeUse),
	}
}
