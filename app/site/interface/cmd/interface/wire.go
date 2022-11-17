//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"flag"
	"os"

	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/interface/internal/biz"
	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/interface/internal/conf"
	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/interface/internal/data"
	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/interface/internal/server"
	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/interface/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, tracesdk *tracesdk.TraceProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
