package websocket

import (
	"wappi/internal/usecases/notification"
)

type Notifier struct {
	hub *Hub
}

func NewNotifier(hub *Hub) *Notifier {
	return &Notifier{hub: hub}
}

func (n *Notifier) NotifyOrderClaimed(payload notification.OrderClaimedPayload) error {
	wsPayload := OrderClaimedPayload{
		OrderID:   payload.OrderID,
		UserID:    payload.UserID,
		ProfileID: payload.ProfileID,
		Status:    payload.Status,
		ETA:       payload.ETA,
		ClaimedAt: payload.ClaimedAt,
	}

	return n.hub.NotifyOrderClaimed(wsPayload)
}

var _ notification.Service = (*Notifier)(nil)
