package appcontext

import (
	"wappi/internal/adapters/datasources"
	"wappi/internal/adapters/datasources/repositories"
	"wappi/internal/platform/config"
)

type Context struct {
	Repositories  *repositories.Repositories
	ConfigService *config.ConfigurationService
}

type Option func(*Context)

type Factory func(opts ...Option) *Context

func NewFactory(
	datasources *datasources.Datasources,
	configService *config.ConfigurationService,
) func(opts ...Option) *Context {
	return func(opts ...Option) *Context {
		return &Context{
			Repositories:  repositories.NewFactory(datasources)(),
			ConfigService: configService,
		}
	}
}
