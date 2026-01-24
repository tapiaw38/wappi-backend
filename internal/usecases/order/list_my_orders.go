package order

import (
	"context"

	"wappi/internal/adapters/datasources/repositories/order"
	"wappi/internal/domain"
	apperrors "wappi/internal/platform/errors"
)

// ListMyOrdersOutputItem represents an order in the user's order list
type ListMyOrdersOutputItem struct {
	ID            string                  `json:"id"`
	ProfileID     *string                 `json:"profile_id,omitempty"`
	Status        string                  `json:"status"`
	StatusMessage *string                 `json:"status_message,omitempty"`
	StatusIndex   int                     `json:"status_index"`
	ETA           string                  `json:"eta"`
	Data          *ListMyOrdersOutputData `json:"data,omitempty"`
	CreatedAt     string                  `json:"created_at"`
	UpdatedAt     string                  `json:"updated_at"`
	AllStatuses   []string                `json:"all_statuses"`
}

// ListMyOrdersOutputData represents the order data in the output
type ListMyOrdersOutputData struct {
	Items []ListMyOrdersOutputItem2 `json:"items"`
}

// ListMyOrdersOutputItem2 represents a single item in the order output
type ListMyOrdersOutputItem2 struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// ListMyOrdersOutput represents the output for listing user's orders
type ListMyOrdersOutput struct {
	Orders []ListMyOrdersOutputItem `json:"orders"`
	Total  int                      `json:"total"`
}

// ListMyOrdersUsecase defines the interface for listing user's orders
type ListMyOrdersUsecase interface {
	Execute(ctx context.Context, userID string) (*ListMyOrdersOutput, apperrors.ApplicationError)
}

type listMyOrdersUsecase struct {
	repo order.Repository
}

// NewListMyOrdersUsecase creates a new instance of ListMyOrdersUsecase
func NewListMyOrdersUsecase(repo order.Repository) ListMyOrdersUsecase {
	return &listMyOrdersUsecase{repo: repo}
}

// Execute lists all orders for a specific user
func (u *listMyOrdersUsecase) Execute(ctx context.Context, userID string) (*ListMyOrdersOutput, apperrors.ApplicationError) {
	orders, err := u.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	allStatuses := make([]string, len(domain.ValidStatuses))
	for i, s := range domain.ValidStatuses {
		allStatuses[i] = string(s)
	}

	output := &ListMyOrdersOutput{
		Orders: make([]ListMyOrdersOutputItem, 0, len(orders)),
		Total:  len(orders),
	}

	for _, o := range orders {
		// Convert domain data to output data
		var outputData *ListMyOrdersOutputData
		if o.Data != nil && len(o.Data.Items) > 0 {
			items := make([]ListMyOrdersOutputItem2, len(o.Data.Items))
			for i, item := range o.Data.Items {
				items[i] = ListMyOrdersOutputItem2{
					Name:     item.Name,
					Price:    item.Price,
					Quantity: item.Quantity,
				}
			}
			outputData = &ListMyOrdersOutputData{Items: items}
		}

		output.Orders = append(output.Orders, ListMyOrdersOutputItem{
			ID:            o.ID,
			ProfileID:     o.ProfileID,
			Status:        string(o.Status),
			StatusMessage: o.StatusMessage,
			StatusIndex:   o.StatusIndex(),
			ETA:           o.ETA,
			Data:          outputData,
			CreatedAt:     o.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:     o.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			AllStatuses:   allStatuses,
		})
	}

	return output, nil
}
