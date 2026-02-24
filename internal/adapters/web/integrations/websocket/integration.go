package websocket

import (
	"yego/internal/platform/config"
	ws "yego/internal/adapters/web/websocket"
)

type Integration interface {
	GetHub() *ws.Hub
}

type integration struct {
	hub *ws.Hub
}

func NewIntegration(cfg *config.ConfigurationService) Integration {
	hub := ws.NewHub()
	go hub.Run()
	return &integration{
		hub: hub,
	}
}

func (i *integration) GetHub() *ws.Hub {
	return i.hub
}
