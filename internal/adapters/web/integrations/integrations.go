package integrations

import (
	"wappi/internal/adapters/web/integrations/auth"
	"wappi/internal/adapters/web/integrations/payments"
	"wappi/internal/adapters/web/integrations/websocket"
	"wappi/internal/platform/config"
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
