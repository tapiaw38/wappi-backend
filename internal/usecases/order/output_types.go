package order

import (
	"wappi/internal/domain"
)

// OrderOutputData represents basic order data for outputs
type OrderOutputData struct {
	ID          string          `json:"id"`
	ProfileID   *string         `json:"profile_id,omitempty"`
	UserID      *string         `json:"user_id,omitempty"`
	Status      string          `json:"status"`
	StatusIndex int             `json:"status_index"`
	ETA         string          `json:"eta"`
	Data        *OrderItemsData `json:"data,omitempty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	AllStatuses []string        `json:"all_statuses,omitempty"`
}

// OrderItemsData represents the items data in an order
type OrderItemsData struct {
	Items []OrderItemOutput `json:"items"`
}

// OrderItemOutput represents a single item in the order output
type OrderItemOutput struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Weight   *int    `json:"weight,omitempty"`
}

// toOrderOutputData converts a domain order to output data
func toOrderOutputData(order *domain.Order, includeStatuses bool) OrderOutputData {
	output := OrderOutputData{
		ID:          order.ID,
		ProfileID:   order.ProfileID,
		UserID:      order.UserID,
		Status:      string(order.Status),
		StatusIndex: order.StatusIndex(),
		ETA:         order.ETA,
		CreatedAt:   order.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   order.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	if order.Data != nil && len(order.Data.Items) > 0 {
		items := make([]OrderItemOutput, len(order.Data.Items))
		for i, item := range order.Data.Items {
			items[i] = OrderItemOutput{
				Name:     item.Name,
				Price:    item.Price,
				Quantity: item.Quantity,
				Weight:   item.Weight,
			}
		}
		output.Data = &OrderItemsData{Items: items}
	}

	if includeStatuses {
		allStatuses := make([]string, len(domain.ValidStatuses))
		for i, s := range domain.ValidStatuses {
			allStatuses[i] = string(s)
		}
		output.AllStatuses = allStatuses
	}

	return output
}

// getAllStatuses returns all valid order statuses
func getAllStatuses() []string {
	allStatuses := make([]string, len(domain.ValidStatuses))
	for i, s := range domain.ValidStatuses {
		allStatuses[i] = string(s)
	}
	return allStatuses
}
