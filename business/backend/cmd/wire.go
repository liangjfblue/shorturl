//go:build wireinject
// +build wireinject

package main

import (
	"backend/internal/component"
	"backend/internal/config"
	"backend/internal/dao"
	"backend/internal/handler"
	"backend/internal/router"
	"backend/internal/server"
	"backend/internal/service"
	"github.com/google/wire"
)

func wireServer() (*server.Server, func(), error) {
	panic(
		wire.Build(
			config.ProviderSet,
			component.ProviderSet,
			dao.ProviderSet,
			service.ProviderSet,
			handler.ProviderSet,
			router.ProviderSet,
			server.NewHttpServer,
			server.NewServer,
		),
	)
}
