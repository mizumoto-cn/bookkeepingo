package data

import (
	accountV1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/internal/conf"

	"context"

	consul "github.com/go-kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
)

// Data .
type Data struct {
	log *log.Helper
	ac  *accountV1.AccountServiceClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	ac *accountV1.AccountServiceClient,
) (*Data, error) {
	log := log.NewHelper(log.With(logger, "module", "site-interface/internal/data"))
	return &Data{
		log: log,
		ac:  ac,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = conf.Consul.Addr
	consulConfig.Scheme = conf.Consul.Scheme
	consulClient, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	return consul.New(consulClient, consul.WithHealthCheck(false))
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = conf.Consul.Addr
	consulConfig.Scheme = conf.Consul.Scheme
	consulClient, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	return consul.New(consulClient, consul.WithHealthCheck(false))
}

func NewAccountServiceClient(
	authConfig *conf.Auth,
	registry registry.Discovery,
	tracer *tracesdk.TracerProvider,
) accountV1.AccountServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///bookkeepingo.mizumoto.tech.account.service"),
		grpc.WithDiscovery(registry),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tracer)),
			recovery.Recovery(),
			jwt.Client(func(token *jwt4.Token) (interface{}, error) {
				return []byte(authConfig.ServiceKey), nil
			}, jwt.WithSigningMethod(jwt4.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	return accountV1.NewAccountServiceClient(conn)
}
