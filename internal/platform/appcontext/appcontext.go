package appcontext

import (
	"wappi/internal/adapters/datasources"
	"wappi/internal/adapters/datasources/repositories"
	"wappi/internal/adapters/web/integrations"
	"wappi/internal/platform/config"
)

type Context struct {
	Repositories  *repositories.Repositories
	Integrations  *integrations.Integrations
	ConfigService *config.ConfigurationService
}

type Option func(*Context)

type Factory func(opts ...Option) *Context

func NewFactory(
	datasources *datasources.Datasources,
	integrations *integrations.Integrations,
	configService *config.ConfigurationService,
) func(opts ...Option) *Context {
	return func(opts ...Option) *Context {
		return &Context{
			Repositories:  repositories.NewFactory(datasources)(),
			Integrations:  integrations,
			ConfigService: configService,
		}
	}
}
