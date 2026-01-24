package admin

import (
	"context"

	orderRepo "wappi/internal/adapters/datasources/repositories/order"
	"wappi/internal/domain"
	apperrors "wappi/internal/platform/errors"
)

// OrderOutput represents an order in the admin list
type OrderOutput struct {
	ID            string          `json:"id"`
	ProfileID     *string         `json:"profile_id,omitempty"`
	UserID        *string         `json:"user_id,omitempty"`
	Status        string          `json:"status"`
	StatusMessage *string         `json:"status_message,omitempty"`
	StatusIndex   int             `json:"status_index"`
	ETA           string          `json:"eta"`
	Data          *OrderDataOutput `json:"data,omitempty"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	AllStatuses   []string        `json:"all_statuses"`
}

// OrderDataOutput represents the order data in the output
type OrderDataOutput struct {
	Items []OrderItemOutput `json:"items"`
}

// OrderItemOutput represents a single item in the order output
type OrderItemOutput struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// ListOrdersOutput represents the output for listing orders
type ListOrdersOutput struct {
	Orders []OrderOutput `json:"orders"`
	Total  int           `json:"total"`
}

// ListOrdersUsecase defines the interface for listing orders
type ListOrdersUsecase interface {
	Execute(ctx context.Context) (*ListOrdersOutput, apperrors.ApplicationError)
}

type listOrdersUsecase struct {
	repo orderRepo.Repository
}

// NewListOrdersUsecase creates a new instance of ListOrdersUsecase
func NewListOrdersUsecase(repo orderRepo.Repository) ListOrdersUsecase {
	return &listOrdersUsecase{repo: repo}
}

// Execute lists all orders
func (u *listOrdersUsecase) Execute(ctx context.Context) (*ListOrdersOutput, apperrors.ApplicationError) {
	orders, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allStatuses := make([]string, len(domain.ValidStatuses))
	for i, s := range domain.ValidStatuses {
		allStatuses[i] = string(s)
	}

	output := &ListOrdersOutput{
		Orders: make([]OrderOutput, 0, len(orders)),
		Total:  len(orders),
	}

	for _, o := range orders {
		// Convert domain data to output data
		var outputData *OrderDataOutput
		if o.Data != nil && len(o.Data.Items) > 0 {
			items := make([]OrderItemOutput, len(o.Data.Items))
			for i, item := range o.Data.Items {
				items[i] = OrderItemOutput{
					Name:     item.Name,
					Price:    item.Price,
					Quantity: item.Quantity,
				}
			}
			outputData = &OrderDataOutput{Items: items}
		}

		output.Orders = append(output.Orders, OrderOutput{
			ID:            o.ID,
			ProfileID:     o.ProfileID,
			UserID:        o.UserID,
			Status:        string(o.Status),
			StatusMessage:  o.StatusMessage,
			StatusIndex:    o.StatusIndex(),
			ETA:            o.ETA,
			Data:           outputData,
			CreatedAt:      o.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:      o.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			AllStatuses:    allStatuses,
		})
	}

	return output, nil
}
