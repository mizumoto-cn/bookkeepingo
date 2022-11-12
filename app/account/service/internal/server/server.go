package server

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2/registry"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	cCfg := consulAPI.DefaultConfig()
	cCfg.Address = conf.Consul.Address
	cCfg.Scheme = conf.Consul.Scheme
	cCli, err := consulAPI.NewClient(cCfg)
	if err != nil {
		panic(err)
	}
	r := consul.New(cCli, consul.WithHealthCheck(false))
	return r
}
