package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	accountV1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewAccountServiceClient,
	NewDiscovery,
	NewRegistrar,
	NewAccountRepo,
)

// Data .
type Data struct {
	log    *log.Helper
	client accountV1.AccountServiceClient
}

// NewData .
func NewData(
	ac accountV1.AccountServiceClient,
	conf *conf.Data,
	logger log.Logger,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "site-admin/data"))
	return &Data{
		log:    l,
		client: ac,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = conf.Consul.Addr
	consulConfig.Scheme = conf.Consul.Scheme
	client, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	return consul.New(client, consul.WithHealthCheck(false))
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = conf.Consul.Addr
	consulConfig.Scheme = conf.Consul.Scheme
	client, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	return consul.New(client, consul.WithHealthCheck(false))
}

func NewAccountServiceClient(rd registry.Discovery, tp *tracesdk.TracerProvider) accountV1.AccountServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///bookkeepingo.mizumoto.tech.account.service"),
		grpc.WithDiscovery(rd),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	ret := accountV1.NewAccountServiceClient(conn)
	return ret
}
