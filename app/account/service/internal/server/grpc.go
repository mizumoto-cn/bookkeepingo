package server

import (
	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/conf"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, au *conf.Auth, logger log.Logger, tp *tracesdk.TracerProvider, s *service.AccountService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(au.JwtSecret), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAccountServiceServer(srv, s)
	return srv
}
