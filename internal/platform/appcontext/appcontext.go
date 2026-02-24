package appcontext

import (
	"yego/internal/adapters/datasources"
	"yego/internal/adapters/datasources/repositories"
	"yego/internal/adapters/web/integrations"
	"yego/internal/platform/config"
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
