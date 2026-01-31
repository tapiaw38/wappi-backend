package admin

import (
	"wappi/internal/domain"
)

// OrderOutput represents an order in the admin list
type OrderOutput struct {
	ID            string            `json:"id"`
	ProfileID     *string           `json:"profile_id,omitempty"`
	UserID        *string           `json:"user_id,omitempty"`
	Status        string            `json:"status"`
	StatusMessage *string           `json:"status_message,omitempty"`
	StatusIndex   int               `json:"status_index"`
	ETA           string            `json:"eta"`
	Data          *domain.OrderData `json:"data,omitempty"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
	AllStatuses   []string          `json:"all_statuses"`
}

// ProfileOutput represents a profile in the admin list
type ProfileOutput struct {
	ID          string          `json:"id"`
	UserID      string          `json:"user_id"`
	PhoneNumber string          `json:"phone_number"`
	Location    *LocationOutput `json:"location,omitempty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

// LocationOutput represents a location in the output
type LocationOutput struct {
	ID        string  `json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
}

// toOrderOutput converts a domain order to output
func toOrderOutput(order *domain.Order) OrderOutput {
	allStatuses := make([]string, len(domain.ValidStatuses))
	for i, s := range domain.ValidStatuses {
		allStatuses[i] = string(s)
	}

	return OrderOutput{
		ID:            order.ID,
		ProfileID:     order.ProfileID,
		UserID:        order.UserID,
		Status:        string(order.Status),
		StatusMessage: order.StatusMessage,
		StatusIndex:   order.StatusIndex(),
		ETA:           order.ETA,
		Data:          order.Data,
		CreatedAt:     order.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     order.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		AllStatuses:   allStatuses,
	}
}
