package domain

import "time"

// Transaction represents a payment transaction in the system
type Transaction struct {
	ID                string    `json:"id"`
	OrderID           string    `json:"order_id"`
	UserID            string    `json:"user_id"`
	ProfileID         *string   `json:"profile_id,omitempty"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	Status            string    `json:"status"`
	PaymentID         *int      `json:"payment_id,omitempty"`
	GatewayPaymentID  *string   `json:"gateway_payment_id,omitempty"`
	CollectorID       *string   `json:"collector_id,omitempty"`
	Description       *string   `json:"description,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
