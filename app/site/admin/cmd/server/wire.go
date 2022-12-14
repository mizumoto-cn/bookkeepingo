//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/biz"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/conf"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/data"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/server"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	wire "github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
