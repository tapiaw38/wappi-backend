package notification

// OrderClaimedPayload contains data about a claimed order
type OrderClaimedPayload struct {
	OrderID   string `json:"order_id"`
	UserID    string `json:"user_id"`
	ProfileID string `json:"profile_id,omitempty"`
	Status    string `json:"status"`
	ETA       string `json:"eta"`
	ClaimedAt string `json:"claimed_at"`
}

// Service defines the interface for sending notifications to clients
// This is a driven port (output port) in hexagonal architecture
type Service interface {
	// NotifyOrderClaimed sends a notification when an order is claimed by a user
	NotifyOrderClaimed(payload OrderClaimedPayload) error
}
