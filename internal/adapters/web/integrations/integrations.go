package integrations

import (
	"yego/internal/adapters/web/integrations/auth"
	"yego/internal/adapters/web/integrations/payments"
	"yego/internal/adapters/web/integrations/websocket"
	"yego/internal/platform/config"
)

type Integrations struct {
	WebSocket websocket.Integration
	Payments  payments.Integration
	Auth      auth.Integration
}

func CreateIntegration(cfg *config.ConfigurationService) *Integrations {
	return &Integrations{
		WebSocket: websocket.NewIntegration(cfg),
		Payments:  payments.NewIntegration(cfg),
		Auth:      auth.NewIntegration(cfg),
	}
}
